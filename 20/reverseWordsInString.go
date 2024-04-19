package main

import (
	"fmt"
	"strings"
)

// reverseWordsInString переворачивает слова в строке
func reverseWordsInString(s string) string {
	words := strings.Fields(s) // Разделяем строку на слова

	// Переворачиваем порядок слов
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// Собираем перевернутые слова обратно в строку
	return strings.Join(words, " ")
}

func main() {
	inputString := "snow dog sun — sun dog snow"

	reversedString := reverseWordsInString(inputString)

	fmt.Println("Перевернутые слова в строке:", reversedString)
}
