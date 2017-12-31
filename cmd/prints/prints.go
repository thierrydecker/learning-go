package main

import (
	"fmt"
)

func main() {

	// Prints all the arguments to console
	// No space between arguments
	// No newline at the end
	fmt.Print("Thierry", "DECKER", "\n")

	// Prints all the arguments to console
	// Space between arguments
	// Newline at the end
	fmt.Println("Thierry", "DECKER")

	// Prints the formatted string to the console
	// No newline at the end
	fmt.Printf("Thierry %s\n", "DECKER")

	// Bytes printed and error code returned
	n, status := fmt.Printf("Thierry %s\n", "DECKER")
	fmt.Printf("%d bytes printed, error code %v returned", n, status)

}
