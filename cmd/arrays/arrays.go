package main

import (
	"fmt"
)

func main() {

	var array01 [5]int
	array01[4] = 25
	fmt.Printf("array01 is                          : %v\n", array01)
	fmt.Printf("The element at position 4 of array01: %d\n", array01[4])
	fmt.Printf("\n")

	var array02 [5]float64
	var average01 float64
	var average02 float64

	array02[0] = 98
	array02[1] = 93
	array02[2] = 77
	array02[3] = 82
	array02[4] = 83

	for i := 0; i < len(array02); i++ {
		average01 += array02[i]
	}

	for _, valueAti := range array02 {
		average02 += valueAti
	}
	average01 = average01 / float64(len(array02))
	average02 = average02 / float64(len(array02))

	fmt.Printf("array02 is                          : %v\n", array02)
	fmt.Printf("Length of array02 is                : %d\n", len(array02))
	fmt.Printf("The average of array02 elements is  : %f\n", average01)
	fmt.Printf("The average of array02 elements is  : %f\n", average02)

}
