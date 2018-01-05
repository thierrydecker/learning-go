package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Total= %d\n", summarize(1, 2, 3))
	fmt.Printf("Total= %d\n", summarize(1, 2, 3, 4, 5))
	numbers := []int64{10, 20, 30}
	fmt.Printf("Total= %d\n", summarize(numbers...))
}

func summarize(numbers ...int64) (total int64) {
	for _, number := range numbers {
		total += number
	}
	return
}
