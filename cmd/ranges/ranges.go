package main

import (
	"fmt"
)

func main() {

	array01 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	array02 := [3][3]int{{11, 12, 13}, {21, 22, 23}, {31, 32, 33},}
	string01 := "Welcome to Go!"

	fmt.Printf("array01: %v\n", array01)
	fmt.Printf("array02: %v\n", array02)

	for position, value := range array01 {
		fmt.Printf("Value at position %d is %d\n", position, value)
	}

	for _, value := range array01 {
		fmt.Printf("Next value of array01 is %d\n", value)
	}

	for position, value01 := range array02 {
		fmt.Printf("Value at position %d is %d\n", position, value01)
		for position, value02 := range value01 {
			fmt.Printf("Value at position %d is %d\n", position, value02)
		}
	}

	for _, value01 := range array02 {
		fmt.Printf("Next value of array01 is %d\n", value01)
		for _, value02 := range value01 {
			fmt.Printf("Next value of array01 is %d\n", value02)
		}
	}

	for _, letter := range string01 {
		fmt.Printf("Next letter of string01 is %c\n", letter)
	}

}
