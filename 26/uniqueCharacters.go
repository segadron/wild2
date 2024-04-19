package main

import (
	"fmt"
	"strings"
)

func checkUnique(s string) bool {
	// Преобразование строки в нижний регистр
	s = strings.ToLower(s)

	// Создание массива символов для отслеживания повторений
	chars := make([]bool, 128)

	// Проверка уникальности символов
	for _, c := range s {
		if chars[c] {
			return false
		}
		chars[c] = true
	}

	return true
}

func main() {
	fmt.Println(checkUnique("abcd"))      // true
	fmt.Println(checkUnique("abCdefAaf")) // false
	fmt.Println(checkUnique("aabcd"))     // false
}
