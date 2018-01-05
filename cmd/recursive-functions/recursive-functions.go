package main

import (
	"fmt"
)

func factorial(x float64) float64 {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func syracuse(x int64) int64 {
	fmt.Printf("%d,", x)
	if x == 0 || x == 1 {
		return 1
	} else if x%2 == 0 {
		return syracuse(x / 2)
	} else {
		return syracuse(3*x + 1)
	}
}

func main() {
	var x float64 = 20
	fmt.Printf("Factorial(%v)= %v\n", x, factorial(x))
	var y int64 = 28
	fmt.Printf("syracuse(%d)= [", y)
	syracuse(y)
	fmt.Printf("\b]")
}
