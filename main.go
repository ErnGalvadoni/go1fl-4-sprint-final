package main

import (
	"fmt"
	"fitness-tracker/daysteps"
	"fitness-tracker/spentcalories"
)

func main() {
	// Тестируем DayActionInfo
	fmt.Println("=== Тестируем DayActionInfo ===")
	data := "678,0h50m"
	info := daysteps.DayActionInfo(data, 70.0, 1.75)
	fmt.Println(info)
	fmt.Println()

	// Тестируем TrainingInfo
	fmt.Println("=== Тестируем TrainingInfo ===")
	trainingData := "3456,Ходьба,3h00m"
	trainingInfo, err := spentcalories.TrainingInfo(trainingData, 70.0, 1.75)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println(trainingInfo)
	}
}