package spentcalories

import (
    "fmt"
    "log"
    "strconv"
    "strings"
    "time"
)

const (
    stepLengthCoefficient      = 0.65
    mInKm                      = 1000
    minInH                     = 60
    walkingCaloriesCoefficient = 0.035
)

func parseTraining(data string) (int, string, time.Duration, error) {
    parts := strings.Split(data, ",")
    
    if len(parts) != 3 {
        return 0, "", 0, fmt.Errorf("неверный формат данных")
    }
    
    stepsStr := strings.TrimSpace(parts[0])
    steps, err := strconv.Atoi(stepsStr)
    if err != nil {
        return 0, "", 0, fmt.Errorf("неверный формат шагов: %v", err)
    }
    
    trainingType := strings.TrimSpace(parts[1])
    
    durationStr := strings.TrimSpace(parts[2])
    duration, err := time.ParseDuration(durationStr)
    if err != nil {
        return 0, "", 0, fmt.Errorf("неверный формат продолжительности: %v", err)
    }
    
    return steps, trainingType, duration, nil
}

func distance(steps int, height float64) float64 {
    stepLength := height * stepLengthCoefficient
    distanceMeters := float64(steps) * stepLength
    distanceKm := distanceMeters / mInKm
    return distanceKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
    if duration <= 0 {
        return 0
    }
    
    dist := distance(steps, height)
    speed := dist / duration.Hours()
    return speed
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
    if steps <= 0 {
        return 0, fmt.Errorf("количество шагов должно быть положительным")
    }
    if weight <= 0 {
        return 0, fmt.Errorf("вес должен быть положительным")
    }
    if height <= 0 {
        return 0, fmt.Errorf("рост должен быть положительным")
    }
    if duration <= 0 {
        return 0, fmt.Errorf("продолжительность должна быть положительной")
    }
    
    speed := meanSpeed(steps, height, duration)
    durationMinutes := duration.Minutes()
    calories := (weight * speed * durationMinutes) / minInH
    
    return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
    if steps <= 0 {
        return 0, fmt.Errorf("количество шагов должно быть положительным")
    }
    if weight <= 0 {
        return 0, fmt.Errorf("вес должен быть положительным")
    }
    if height <= 0 {
        return 0, fmt.Errorf("рост должен быть положительным")
    }
    if duration <= 0 {
        return 0, fmt.Errorf("продолжительность должна быть положительной")
    }
    
    speed := meanSpeed(steps, height, duration)
    durationMinutes := duration.Minutes()
    calories := (weight * speed * durationMinutes) / minInH
    calories = calories * walkingCaloriesCoefficient
    
    return calories, nil
}

func TrainingInfo(data string, weight, height float64) (string, error) {
    steps, trainingType, duration, err := parseTraining(data)
    if err != nil {
        log.Println(err)
        return "", err
    }
    
    var calories float64
    var dist float64
    var speed float64
    
    switch trainingType {
    case "Ходьба":
        calories, err = WalkingSpentCalories(steps, weight, height, duration)
        if err != nil {
            return "", err
        }
        dist = distance(steps, height)
        speed = meanSpeed(steps, height, duration)
        
    case "Бег":
        calories, err = RunningSpentCalories(steps, weight, height, duration)
        if err != nil {
            return "", err
        }
        dist = distance(steps, height)
        speed = meanSpeed(steps, height, duration)
        
    default:
        return "", fmt.Errorf("неизвестный тип тренировки: %s", trainingType)
    }
    
    result := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f",
        trainingType,
        duration.Hours(),
        dist,
        speed,
        calories)
    
    return result, nil
}
