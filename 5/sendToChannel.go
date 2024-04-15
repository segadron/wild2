package main

import (
	"time"
)

func main() {
	channel := make(chan int)

	go func() {
		for i := 0; ; i++ {
			channel <- i
			time.Sleep(time.Second)
		}
	}()

	go func() {
		time.Sleep(5 * time.Second)
		close(channel)
	}()

	for value := range channel {
		println(value)
	}
}
