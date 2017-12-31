package main

import (
	"fmt"
)

func main() {

	const myConstant float64 = 3.14159
	fmt.Printf("myConstant= %f", myConstant)

	// This line will fail at compilation time
	myConstant = 3.14

}
