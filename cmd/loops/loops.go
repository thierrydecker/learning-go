package main

import (
	"fmt"
)

func main() {

	// Classical for loop
	for i := 0; i < 10; i++ {
		fmt.Printf("i= %d\n", i)
	}

	// While loop
	j := 0
	for j < 10 {
		fmt.Printf("j= %d\n", j)
		j++
	}

}
