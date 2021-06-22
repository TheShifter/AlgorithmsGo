package main

import (
	"math/rand"
)

func QuickSort(list []int) []int {
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
	left = QuickSort(left)
	right = QuickSort(right)
	return append(append(left, []int{v}...), right...)
}
