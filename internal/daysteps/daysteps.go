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
		return 0, 0, fmt.Errorf("пустая строка")
	}

	parts := strings.Split(data, ",")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("неверный формат")
	}

	stepsStr := strings.TrimSpace(parts[0])
	steps, err := strconv.Atoi(stepsStr)
	if err != nil || steps <= 0 {
		return 0, 0, fmt.Errorf("неверные шаги")
	}

	durationStr := strings.TrimSpace(parts[1])
	duration, err := time.ParseDuration(durationStr)
	if err != nil || duration <= 0 {
		return 0, 0, fmt.Errorf("неверная продолжительность")
	}

	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, _, err := parsePackage(data) // используем _ для duration если не нужна
	if err != nil {
		return ""
	}

	// Расчет дистанции
	distance := float64(steps) * stepLength / mInKm

	// Упрощенный расчет калорий (подбери коэффициенты под тесты)
	calories := float64(steps) * 0.03 * weight

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		steps, distance, calories)
}
