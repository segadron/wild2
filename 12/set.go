package main

import "fmt"

func main() {
	// Последовательность строк
	strings := []string{"cat", "cat", "dog", "cat", "tree"}

	// Создаем пустой словарь
	uniqueStrings := make(map[string]bool)

	// Добавляем строки в словарь
	for _, str := range strings {
		uniqueStrings[str] = true
	}

	// Выводим уникальные строки
	for str := range uniqueStrings {
		fmt.Println(str)
	}
}
