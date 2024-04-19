package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Counter - структура-счетчик
type Counter struct {
	count int64 // поле для хранения значения счетчика
}

// NewCounter - функция для создания нового счетчика
func NewCounter() *Counter {
	return &Counter{}
}

// Increment - метод для увеличения значения счетчика
func (c *Counter) Increment() {
	atomic.AddInt64(&c.count, 1) // атомарная операция увеличения значения счетчика на 1
}

// GetCount - метод для получения текущего значения счетчика
func (c *Counter) GetCount() int64 {
	return atomic.LoadInt64(&c.count) // атомарная операция чтения значения счетчика
}

func main() {
	counter := NewCounter() // создаем новый счетчик
	var wg sync.WaitGroup   // группа синхронизации для ожидания завершения горутин
	wg.Add(10)

	// создаем 10 горутин, каждая из которых инкрементирует счетчик 1000 раз
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Increment()
			}
		}()
	}

	wg.Wait()                                       // ожидаем завершения всех горутин
	fmt.Println("Final count:", counter.GetCount()) // выводим итоговое значение счетчика
}
