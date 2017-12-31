package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Printf("%d is even and i*i = %d", i, i*i)
		} else {
			fmt.Printf("%d is odd", i)
		}
		if i > 5 {
			fmt.Printf(". (%d is greater than 5)", i)
		}
		fmt.Println()
	}
}
