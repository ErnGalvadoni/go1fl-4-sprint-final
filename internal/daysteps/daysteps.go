package daysteps

import (
    "fmt"
    "strconv"
    "strings"
    "time"
)

const (
    stepLength = 0.65 // средняя длина шага в метрах
    mInKm      = 1000 // метров в километре
)

func parsePackage(data string) (int, time.Duration, error) {
    parts := strings.Split(data, ",")
    
    if len(parts) != 2 {
        return 0, 0, fmt.Errorf("incorrect data format")
    }
    
    stepsStr := strings.TrimSpace(parts[0])
    steps, err := strconv.Atoi(stepsStr)
    if err != nil {
        return 0, 0, fmt.Errorf("неверный формат шагов: %w", err)
    }
    
    if steps <= 0 {
        return 0, 0, fmt.Errorf("количество шагов должно быть положительным")
    }
    
    durationStr := strings.TrimSpace(parts[1])
    duration, err := time.ParseDuration(durationStr)
    if err != nil {
        return 0, 0, fmt.Errorf("неверный формат продолжительности: %v", err)
    }
    
    return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
    steps, duration, err := parsePackage(data)
    if err != nil {
        fmt.Println("Ошибка:", err)
        return ""
    }
    
    if steps <= 0 {
        return ""
    }
    
    distanceMeters := float64(steps) * stepLength
    distanceKm := distanceMeters / mInKm
    
    calories, err := WalkingSpentCalories(steps, weight, height, duration)
    if err != nil {
        fmt.Println("Ошибка расчета калорий:", err)
        return ""
    }
    
    result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.", 
        steps, distanceKm, calories)
    
    return result
}
