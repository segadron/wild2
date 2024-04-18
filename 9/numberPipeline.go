package main

import (
	"fmt"
)

func main() {
	input := []int{1, 2, 3, 4, 5}

	// Создаем каналы
	ch1 := make(chan int)
	ch2 := make(chan int)

	// Записываем числа из массива в первый канал
	go func() {
		for _, num := range input {
			ch1 <- num
		}
		close(ch1)
	}()

	// Умножаем числа на 2 и пишем результат во второй канал
	go func() {
		for num := range ch1 {
			ch2 <- num * 2
		}
		close(ch2)
	}()

	// Выводим данные из второго канала в stdout
	for result := range ch2 {
		fmt.Println(result)
	}
}
