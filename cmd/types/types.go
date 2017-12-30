package main

import (
	"fmt"
)

func main() {

	// Strings

	// Declaration then assignment
	var string01 string
	string01 = "First string"
	fmt.Printf("string01                         = %s\n", string01)

	// Declaration then assignment
	var string02 string = "Second string"
	fmt.Printf("string02                         = %s\n", string02)

	// Implicit declaration and assignment
	string03 := "Third string"
	fmt.Printf("string03                         = %s\n", string03)

	// Integers

	var integer01 int8 = 127
	fmt.Printf("integer01 int8 (-128 to 127)     = %d\n", integer01)
	var integer02 uint8 = 255
	fmt.Printf("integer02 uint8 (0 to 255)       = %d\n", integer02)
	var integer03 int16 = 32767
	fmt.Printf("integer04 int16 (-32768 to 32767)= %d\n", integer03)
	var integer04 uint16 = 65535
	fmt.Printf("integer04 uint16 (0 to 65535)    = %d\n", integer04)

}
