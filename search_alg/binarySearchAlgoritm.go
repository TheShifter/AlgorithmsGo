package main

import "fmt"

func main() {
	values := make([]int, 100)
	for i := 0; i < 100; i++ {
		values[i] = i
	}
	fmt.Println(values)
	searchNumber := binarySearch(values, 93)
	fmt.Println(searchNumber)
}

func binarySearch(values []int, number int) int {
	var low int
	high := len(values) - 1

	for low <= high {
		median := (low + high) / 2
		if values[median] < number {
			low = median + 1
		} else {
			high = median - 1
		}
	}
	if low == len(values) || values[low] != number {
		return -1
	} else {
		return low
	}

}
