package main

import (
	"context" // пакет для работы с контекстами
	"fmt"     // пакет для форматирования строк
	"sync"    // пакет для работы с синхронизацией
	"time"    // пакет для работы с временем
)

func main() {
	// Группа синхронизации для ожидания завершения всех горутин
	var wg sync.WaitGroup
	// Добавляем 3 горутины в группу синхронизации
	wg.Add(3)

	// Горутина с использованием return
	go func() {
		defer wg.Done() // Уменьшаем счетчик горутинов в группе синхронизации
		fmt.Println("Горутина с использованием `return`")
		return // Возвращаемся из горутины
	}()

	// Горутина с использованием panic
	go func() {
		defer wg.Done() // Уменьшаем счетчик горутинов в группе синхронизации
		fmt.Println("Горутина с использованием `panic`")
		panic("panic") // Вызываем паник
	}()

	// Горутина с использованием дефер для восстановления после паники
	go func() {
		defer func() {
			if err := recover(); err != nil { // Принимаем значение, возвращаемое дефером recover
				fmt.Printf("Восстановлено: %v\n", err) // Выводим восстановленное значение
			}
			wg.Done() // Уменьшаем счетчик горутинов в группе синхронизации
		}()

		fmt.Println("Горутина с использованием `defer recover()`")
		panic("panic") // Вызываем паник
	}()

	// Горутина с использованием контекста и WithCancel
	go func() {
		defer wg.Done() // Уменьшаем счетчик горутинов в группе синхронизации
		fmt.Println("Горутина с использованием `context.WithCancel`")
		ctx, cancel := context.WithCancel(context.Background()) // Создаем контекст с помощью WithCancel
		defer cancel()                                          // Отменяем контекст при выходе из горутины

		select {
		case <-time.After(3 * time.Second): // Ожидаем 3 секунды
			fmt.Println("После 3 секунд")
		case <-ctx.Done(): // Ожидаем остановки контекста
			fmt.Println("Контекст остановлен")
		}
	}()

	// Ожидаем завершения всех горутинов в группе синхронизации
	wg.Wait()
	fmt.Println("Программа закончила")
}
