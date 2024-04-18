package main

import "fmt"

// Функция для нахождения пересечения двух неупорядоченных множеств
func intersection(set1, set2 map[int]bool) map[int]bool {
	// Создаем пустое неупорядоченное множество для хранения пересечения
	intersection := make(map[int]bool)

	// Проходим по одному из множеств
	for key := range set1 {
		// Проверяем, присутствует ли текущий ключ в другом множестве
		if set2[key] {
			// Если ключ присутствует, добавляем его в пересечение
			intersection[key] = true
		}
	}

	// Возвращаем пересечение
	return intersection
}

func main() {
	// Примеры неупорядоченных множеств
	set1 := map[int]bool{1: true, 2: true, 3: true}
	set2 := map[int]bool{2: true, 3: true, 4: true}

	// Находим пересечение множеств
	result := intersection(set1, set2)

	// Выводим результат
	fmt.Println(result)
}
