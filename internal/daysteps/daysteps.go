package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	stepLength = 0.65
	mInKm      = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	if data == "" {
		return 0, 0, fmt.Errorf("invalid data format")
	}

	parts := strings.Split(data, ",")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("invalid data format")
	}

	stepsStr := strings.TrimSpace(parts[0])
	steps, err := strconv.Atoi(stepsStr)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid steps format")
	}

	if steps <= 0 {
		return 0, 0, fmt.Errorf("steps must be positive")
	}

	durationStr := strings.TrimSpace(parts[1])
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid duration format")
	}

	if duration <= 0 {
		return 0, 0, fmt.Errorf("duration must be positive")
	}

	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return ""
	}

	// Вычисляем дистанцию
	distanceMeters := float64(steps) * stepLength
	distanceKm := distanceMeters / mInKm

	// УДАЛИТЕ переменную duration если она не используется
	// Формула калорий должна быть правильной
	calories := distanceKm * weight * 0.75

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.", 
		steps, distanceKm, calories)
}
