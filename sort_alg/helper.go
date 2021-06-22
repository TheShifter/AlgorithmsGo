package main

import (
	"math/rand"
	"time"
)

func GenerateSlice(size, randLim int) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(randLim) - rand.Intn(randLim)
	}
	return slice
}
