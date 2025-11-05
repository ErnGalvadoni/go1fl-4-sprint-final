package daysteps

import (
	"fmt"
	"time"
)

const (
	stepLength = 0.65
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// Простая заглушка
	return 1000, time.Hour, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
		return ""
	}
	
	// Используем переменные чтобы избежать ошибок "unused"
	_ = duration
	_ = weight
	_ = height
	
	return fmt.Sprintf("Количество шагов: %d.", steps)
}
