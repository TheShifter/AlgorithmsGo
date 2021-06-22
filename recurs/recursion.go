package main

import "fmt"

func main() {
	countdown(5)
	fmt.Println(factorial(4))
}

func countdown(i int) int {
	fmt.Println(i)
	if i <= 0 {
		return 0
	} else {
		countdown(i - 1)
	}
	return i
}

func factorial(x int) int {
	if x == 1 {
		return 1
	} else {
		return x * factorial(x-1)
	}
}
