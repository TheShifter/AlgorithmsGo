package main

import (
	"fmt"
)

func main() {
	fmt.Println("------------- Quick Sort -------------")
	quickSortSlice := GenerateSlice(100, 100)
	fmt.Println(QuickSort(quickSortSlice))

	fmt.Println("------------- Selection Sort -------------")
	selectionSortSlice := GenerateSlice(9999, 999)
	fmt.Println("\n--- Unsorted --- \n\n", selectionSortSlice)
	SelectionSort(selectionSortSlice)
	fmt.Println("\n--- Sorted --- \n\n", selectionSortSlice)
}
