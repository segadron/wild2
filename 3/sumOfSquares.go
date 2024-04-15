package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup
	wg.Add(len(numbers))

	sumSquares := make(chan int)

	for _, num := range numbers {
		go func(n int) {
			defer wg.Done()
			sumSquares <- n * n
		}(num)
	}

	go func() {
		wg.Wait()
		close(sumSquares)
	}()

	total := 0
	for square := range sumSquares {
		total += square
	}

	fmt.Println("Сумма квадратов:", total)
}
