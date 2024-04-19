package main

import (
	"fmt"
	"time"
)

func sleep(seconds time.Duration) {
	time.Sleep(seconds * time.Second)
	fmt.Println("Slept for", seconds, "seconds")
}

func main() {
	// Пример использования
	sleep(5) // Подождите 5 секунд
}
