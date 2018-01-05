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

func main() {
	var x float64 = 20
	fmt.Printf("Factorial(%v) = %v\n", x, factorial(x))
}
