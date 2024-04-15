package main

import (
	"fmt"
	"sync"
)

func main() {
	// Создаем map для хранения данных
	data := make(map[string]string)
	// Создаем мьютекс для безопасного доступа к map из нескольких горутин
	var mutex sync.Mutex

	var wg sync.WaitGroup
	numRoutines := 10
	wg.Add(numRoutines)

	for i := 0; i < numRoutines; i++ {
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", i)
			value := fmt.Sprintf("value%d", i)

			// Захватываем мьютекс перед записью в map
			mutex.Lock()
			data[key] = value
			// Освобождаем мьютекс после записи
			mutex.Unlock()
		}(i)
	}

	wg.Wait()

	// Выводим данные после завершения всех горутин
	fmt.Println(data)
}
