package main

import "fmt"

// reverseString переворачивает строку
func reverseString(s string) string {
	// Преобразуем строку в массив rune (символы)
	runes := []rune(s)

	// Переворачиваем массив rune
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// Преобразуем массив rune обратно в строку
	return string(runes)
}

func main() {
	// Запрашиваем у пользователя ввод строки
	inputString := "главрыба — абырвалг"

	// Вызываем функцию reverseString для переворота строки
	reversedString := reverseString(inputString)

	// Выводим результат на экран
	fmt.Println("Перевернутая строка:", reversedString)
}
