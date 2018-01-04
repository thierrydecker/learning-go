package main

import (
	"fmt"
)

func fa() {
	fmt.Printf("No input, no ouput, just printing this message\n")
}

func fb(a int16, b int16) {
	fmt.Printf("Only inputs, Sum equal %d\n", a+b)
}

func fc() int16 {
	fmt.Printf("Only output... ")
	return 100
}

func fd() (int16, string) {
	fmt.Printf("Only outputs... ")
	return 100, "This is a string"
}

func fe(a int32, b int32) (mult int32) {
	mult = a * b
	fmt.Printf("Named returns... ")
	return
}

func main() {
	fa()
	fb(1, 2)
	fmt.Printf("%d\n", fc())
	a, b := fd()
	fmt.Printf("%d, %s\n", a, b)
	c := fe(100, 10)
	fmt.Printf("%d\n", c)
}
