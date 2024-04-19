package main

import "fmt"

func removeItemFromSlice(sl []int, i int) []int {
	// Проверяем, что индекс находится в допустимых границах слайса
	if i < 0 || i >= len(sl) {
		return sl
	}

	// Используем оператор append для создания нового слайса,
	// который содержит все элементы, кроме удаляемого
	return append(sl[:i], sl[i+1:]...)
}

func main() {
	// Создаем слайс с некоторыми элементами
	sl := []int{1, 2, 3, 4, 5}

	// Удаляем элемент с индексом 2
	sl = removeItemFromSlice(sl, 2)

	// Выводим обновленный слайс
	fmt.Println(sl)
}
