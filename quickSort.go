package main

import (
	"fmt"
	"math/rand"
	"time"
)

func quickSort(list []int) []int {
	if len(list) < 2 {
		return list
	}

	item := rand.Intn(len(list))

	var left []int
	var right []int

	v := list[item]

	list = append(list[:item])

	for i := 0; i < len(list); i++ {
		if list[i] < v {
			left = append(left, list[i])
		} else {
			right = append(right, list[i])
		}
	}
	left = quickSort(left)
	right = quickSort(right)
	return append(append(left, []int{v}...), right...)
}

func generateSlice(size int) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(100) - rand.Intn(100)
	}
	return slice
}

func main() {
	slice := generateSlice(100)
	fmt.Println(quickSort(slice))
}
