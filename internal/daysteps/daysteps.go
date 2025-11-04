package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// Разделить строку на слайс строк
	parts := strings.Split(data, ",")
	
	// Проверить, чтобы длина слайса была равна 2
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("invalid data format")
	}

	// Преобразовать первый элемент (количество шагов) в int
	stepsStr := strings.TrimSpace(parts[0])
	steps, err := strconv.Atoi(stepsStr)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid steps format")
	}

	// Проверить: количество шагов должно быть больше 0
	if steps <= 0 {
		return 0, 0, fmt.Errorf("steps must be positive")
	}

	// Преобразовать второй элемент в time.Duration
	durationStr := strings.TrimSpace(parts[1])
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid duration format")
	}

	// Вернуть количество шагов, продолжительность и nil
	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// Получить данные с помощью parsePackage()
	steps, duration, err := parsePackage(data)
	if err != nil {
		// В случае ошибки вернуть пустую строку
		return ""
	}

	// Проверить, чтобы количество шагов было больше 0
	if steps <= 0 {
		return ""
	}

	// Вычислить дистанцию в метрах
	distanceMeters := float64(steps) * stepLength
	
	// Перевести дистанцию в километры
	distanceKm := distanceMeters / mInKm

	// Вычислить количество калорий
	// Временно используем упрощенную формулу, так как spentcalories пакет еще не реализован
	calories := calculateCalories(steps, weight, height, duration)

	// Сформировать строку для возврата
	info := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.", 
		steps, distanceKm, calories)

	return info
}

// Вспомогательная функция для расчета калорий
// Замените на вызов WalkingSpentCalories из пакета spentcalories когда он будет готов
func calculateCalories(steps int, weight, height float64, duration time.Duration) float64 {
	// Упрощенная формула расчета калорий
	// MET (Metabolic Equivalent) для ходьбы = 3.5
	met := 3.5
	calories := met * weight * duration.Hours()
	return calories
}
