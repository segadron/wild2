package main

import (
	"fmt"
	"sort"
)

func binarySearch(arr []int, target int) int {
	// Сортируем массив, если он не отсортирован
	sort.Ints(arr)

	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := (left + right) / 2

		if arr[mid] == target {
			return mid
		}

		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 5

	index := binarySearch(arr, target)

	if index == -1 {
		fmt.Println("Элемент не найден в массиве")
	} else {
		fmt.Printf("Элемент %d найден по индексу %d\n", target, index)
	}
}
