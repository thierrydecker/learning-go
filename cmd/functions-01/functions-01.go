package main

import (
	"fmt"
)

func average(elements []float64) float64 {
	total := 0.0
	for _, v := range elements {
		total += v
	}
	return total / float64(len(elements))
}

func main() {
	elementsToAdd := []float64{10, 20, 11, 12, 35}
	fmt.Println(average(elementsToAdd))
}
