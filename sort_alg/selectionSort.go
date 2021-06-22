package main

func SelectionSort(list []int) []int {
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
