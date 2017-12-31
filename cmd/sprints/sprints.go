package main

import (
	"fmt"
)

func main() {

	// Sends all the arguments to a string
	// No space between arguments
	// No newline at the end
	var a string = fmt.Sprint("Thierry", "DECKER", "\n")
	fmt.Print(a)

	// Sends all the arguments to a string
	// Space between arguments
	// Newline at the end
	var b string = fmt.Sprintln("Thierry", "DECKER")
	fmt.Print(b)

	// Sends the formatted string to a string
	// No newline at the end
	var c string = fmt.Sprintf("Thierry %s\n", "DECKER")
	fmt.Print(c)

}
