package main

import (
	"fmt"
	"reflect"
)

func main() {
	var value interface{}

	value = 42
	fmt.Println("Type of value:", reflect.TypeOf(value))

	value = "hello"
	fmt.Println("Type of value:", reflect.TypeOf(value))

	value = true
	fmt.Println("Type of value:", reflect.TypeOf(value))

	value = make(chan int)
	fmt.Println("Type of value:", reflect.TypeOf(value))
}
