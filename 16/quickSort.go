package main

import (
	"fmt"
	"sort"
)

func main() {
	// Создаем массив для сортировки
	numbers := []int{5, 2, 8, 1, 9, 3}

	// Сортируем массив с помощью функции sort.Ints()
	sort.Ints(numbers)

	// Выводим отсортированный массив
	fmt.Println(numbers)
}
