package main

import (
	"fmt"
)

func setToZero(x int8) {
	x = 0
	fmt.Printf("In setToZero(), x is: %d\n", x)
}

func setToOne(xPtr *int8) {
	*xPtr = 1
	fmt.Printf("In setToOne(), x is : %d\n", *xPtr)
}

func main() {
	var x int8 = 10
	fmt.Printf("In main(), x is     : %d\n", x)
	setToZero(x)
	fmt.Printf("In main(), x is     : %d\n", x)
	setToOne(&x)
	fmt.Printf("In main(), x is     : %d\n", x)
}
