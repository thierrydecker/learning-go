package main

import (
	"fmt"
)

var string01 = "string01"

func main() {

	var string02 = "string02"
	fmt.Println(string01)
	fmt.Println(string02)

	f()

}

func f() {

	fmt.Println(string01)
	fmt.Println(string02)

}
