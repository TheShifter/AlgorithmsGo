package main

import (
	"fmt"
	"math/rand"
	"time"
)

func selectionSort(list []int) []int {
	for i := 0; i < len(list); i++ {
		minIndex := i
		for j := minIndex + 1; j < len(list); j++ {
			if list[j] < list[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			oldMinValue := list[i]
			list[i] = list[minIndex]
			list[minIndex] = oldMinValue
		}
	}
	return list
}

func generateSlice(size int) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func main() {
	slice := generateSlice(9999)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	selectionSort(slice)
	fmt.Println("\n--- Sorted --- \n\n", slice)
}
