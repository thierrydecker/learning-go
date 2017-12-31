package main

import (
	"fmt"
)

func main() {

	defer fmt.Printf("End of count down!\n")

	for i := 0; i < 10; i++ {
		defer fmt.Printf("Count: %d\n", i)
	}

}
