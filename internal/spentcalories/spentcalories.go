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

func ParseTraining(data string) (int, time.Duration, error) {
	parts := strings.Split(data, ";")
	if len(parts) != 3 {
		return 0, 0, fmt.Errorf("invalid data format")
	}

	// Парсим количество шагов
	steps, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil {
		return 0, 0, fmt.Errorf("invalid steps format")
	}
	if steps <= 0 {
		return 0, 0, fmt.Errorf("количество шагов должно быть положительным")
	}

	// Парсим продолжительность
	durationStr := strings.TrimSpace(parts[2])
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid duration format")
	}
	if duration <= 0 {
		return 0, 0, fmt.Errorf("продолжительность должна быть положительной")
	}

	return steps, duration, nil
}

func Distance(steps int, height float64) float64 {
	// Используем либо фиксированную длину шага, либо рассчитываем из роста
	var stepLength float64
	if height > 0 {
		stepLength = height * stepLengthCoefficient / 100 // переводим в метры
	} else {
		stepLength = lenStep
	}
	
	distance := float64(steps) * stepLength / mInKm // переводим в километры
	return distance
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	dist := Distance(steps, height)
	hours := duration.Hours()
	
	if hours == 0 {
		return 0
	}
	
	return dist / hours
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	steps, duration, err := ParseTraining(data)
	if err != nil {
		return "", err
	}

	// Парсим тип тренировки
	parts := strings.Split(data, ";")
	if len(parts) != 3 {
		return "", fmt.Errorf("invalid data format")
	}
	trainingType := strings.TrimSpace(parts[0])

	// Проверяем вес и рост
	if weight <= 0 {
		return "", fmt.Errorf("вес должен быть положительным")
	}
	if height <= 0 {
		return "", fmt.Errorf("рост должен быть положительным")
	}

	// Рассчитываем показатели
	dist := Distance(steps, height)
	speed := MeanSpeed(steps, height, duration)
	
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

	// Форматируем результат
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

	speed := MeanSpeed(steps, height, duration)
	
	// Формула для бега
	calories := (18*speed - 20) * weight / mInKm * duration.Hours() * minInH
	
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

	// Формула для ходьбы
	distance := Distance(steps, height)
	calories := (0.035 * weight + (distance/height) * 0.029 * weight) * duration.Hours() * minInH
	
	return calories, nil
}
