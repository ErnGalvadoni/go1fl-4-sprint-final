package daysteps

import (
	"fmt"
	"time"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// Временная заглушка - замени на реальную логику
	if data == "" {
		return 0, 0, fmt.Errorf("пустая строка")
	}
	
	// TODO: реализовать парсинг строки формата "1000,1h30m"
	// Парсинг шагов и продолжительности
	
	return 1000, time.Hour, nil // ← ВРЕМЕННЫЙ RETURN
}

func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
		// Логируем ошибку, если нужно
		return "" // ← RETURN при ошибке
	}
	
	// TODO: реализовать расчет дистанции и калорий
	// Используй константы stepLength и mInKm
	
	return fmt.Sprintf("Шаги: %d, Вес: %.1f, Рост: %.2f", steps, weight, height) // ← ОСНОВНОЙ RETURN
}
