package spentcalories

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	parts := strings.Split(data, ";")
	if len(parts) != 3 {
		return 0, "", 0, fmt.Errorf("invalid data format")
	}

	trainingType := strings.TrimSpace(parts[0])

	steps, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil {
		return 0, "", 0, fmt.Errorf("invalid steps format")
	}
	if steps <= 0 {
		return 0, "", 0, fmt.Errorf("количество шагов должно быть положительным")
	}

	durationStr := strings.TrimSpace(parts[2])
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return 0, "", 0, fmt.Errorf("invalid duration format")
	}
	if duration <= 0 {
		return 0, "", 0, fmt.Errorf("продолжительность должна быть положительной")
	}

	return steps, trainingType, duration, nil
}

func distance(steps int, height float64) float64 {
	// ИСПРАВЛЕНИЕ: убрали деление на 100 для роста
	var stepLength float64
	if height > 0 {
		stepLength = height * stepLengthCoefficient // рост уже в метрах
	} else {
		stepLength = lenStep
	}
	// УМНОЖАЕМ на 1000 чтобы получить метры, потом делим на 1000 для км
	distanceMeters := float64(steps) * stepLength
	return distanceMeters / mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	dist := distance(steps, height)
	hours := duration.Hours()
	if hours == 0 {
		return 0
	}
	return dist / hours
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	steps, trainingType, duration, err := parseTraining(data)
	if err != nil {
		return "", err
	}

	if weight <= 0 {
		return "", fmt.Errorf("вес должен быть положительным")
	}
	if height <= 0 {
		return "", fmt.Errorf("рост должен быть положительным")
	}

	dist := distance(steps, height)
	speed := meanSpeed(steps, height, duration)

	var calories float64
	var caloriesErr error

	switch strings.ToLower(trainingType) {
	case "ходьба", "walking":
		calories, caloriesErr = WalkingSpentCalories(steps, weight, height, duration)
	case "бег", "running":
		calories, caloriesErr = RunningSpentCalories(steps, weight, height, duration)
	default:
		return "", fmt.Errorf("неизвестный тип тренировки: %s", trainingType)
	}

	if caloriesErr != nil {
		return "", caloriesErr
	}

	info := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f",
		trainingType, duration.Hours(), dist, speed, calories)

	return info, nil
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
	
	// ИСПРАВЛЕННАЯ формула для бега
	calories := (18*speed - 20) * weight * duration.Hours()
	
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
	
	// ИСПРАВЛЕННАЯ формула для ходьбы с walkingCaloriesCoefficient
	calories := (weight * 2 * speed) * durationMinutes / minInH * walkingCaloriesCoefficient
	
	return calories, nil
}
