package main

import (
	"fmt"
)

func main() {

	var01 := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var02 := var01[2:5]

	fmt.Printf("var01 = %v, length; %d, capacity: %d\n", var01, len(var01), cap(var01))
	fmt.Printf("var02 = %v, length; %d, capacity: %d\n", var02, len(var02), cap(var02))
	fmt.Printf("\n")

	var02[1] *= 10
	fmt.Printf("var01 = %v, length; %d, capacity: %d\n", var01, len(var01), cap(var01))
	fmt.Printf("var02 = %v, length; %d, capacity: %d\n", var02, len(var02), cap(var02))
	fmt.Printf("\n")

	var02 = var02[:cap(var02)]
	fmt.Printf("var01 = %v, length; %d, capacity: %d\n", var01, len(var01), cap(var01))
	fmt.Printf("var02 = %v, length; %d, capacity: %d\n", var02, len(var02), cap(var02))
	fmt.Printf("\n")

	var02 = var01[:]
	fmt.Printf("var01 = %v, length; %d, capacity: %d\n", var01, len(var01), cap(var01))
	fmt.Printf("var02 = %v, length; %d, capacity: %d\n", var02, len(var02), cap(var02))
	fmt.Printf("\n")

}
