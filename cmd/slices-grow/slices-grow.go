package main

import (
	"fmt"
)

func main() {

	slice01 := make([]byte, 1, 1)
	slice01[0] = 1
	fmt.Printf("slice01: %v, length: %d, capacity: %d\n", slice01, len(slice01), cap(slice01))

	slice02 := make([]byte, len(slice01), cap(slice01)*2)
	for i := range slice01 {
		slice02[i] = slice01[i]
	}
	slice01 = slice02
	fmt.Printf("slice02: %v, length: %d, capacity: %d\n", slice02, len(slice02), cap(slice02))

	slice03 := make([]byte, len(slice02), cap(slice02)*2)
	copy(slice03, slice02)
	fmt.Printf("slice03: %v, length: %d, capacity: %d\n", slice03, len(slice03), cap(slice03))

}
