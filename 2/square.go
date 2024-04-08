package main

import (
	"fmt"
	"sync"
)

// calcSquare вычисляет квадрат числа
// num - число которое нужно вычислить
// result - канал, куда будет отправлен результат
// wg - WaitGroup для ожидания завершения всех задач
func calcSquare(num int, result chan<- int, wg *sync.WaitGroup) {
	// Уменьшение счетчика WaitGroup
	defer wg.Done()
	// Отправка результата в канал
	result <- num * num
}

// main - входная точка программы
func main() {
	// Список чисел для расчета квадрата
	numbers := []int{2, 4, 6, 8, 10}
	// Канал для отправки результатов расчета
	result := make(chan int, len(numbers))
	// WaitGroup для ожидания завершения всех задач
	wg := &sync.WaitGroup{}

	// Расчет квадрата для каждого числа в списке
	for _, num := range numbers {
		wg.Add(1)
		go calcSquare(num, result, wg)
	}

	// Завершение всех задач в WaitGroup и закрытие канала
	go func() {
		wg.Wait()
		close(result)
	}()

	// Вывод результатов расчета в консоль
	for r := range result {
		fmt.Println(r)
	}
}
