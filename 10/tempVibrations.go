package main

import (
	"fmt"
	"math"
)

// groupTemperatures группирует температурные значения с шагом 10 градусов.
// Вход: последовательность температурных значений.
// Выход: карта, где ключ - группа температурных значений, а значение - срез значений в этой группе.
func groupTemperatures(temperatures []float64) map[int][]float64 {
	groups := make(map[int][]float64)

	// Перебираем каждое температурное значение.
	for _, temp := range temperatures {
		// Вычисляем группу температурного значения, округляя его до ближайшего целого значения с шагом 10 градусов.
		group := int(math.Floor(temp/10) * 10)
		// Добавляем температурное значение в соответствующую группу.
		groups[group] = append(groups[group], temp)
	}

	return groups
}

func main() {
	// Создаем пример последовательности температур.
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	// Вызываем функцию groupTemperatures для группировки температурных значений.
	groupedTemperatures := groupTemperatures(temperatures)

	// Выводим группы температурных значений.
	for group, temps := range groupedTemperatures {
		fmt.Printf("%d: %v\n", group, temps)
	}
}
