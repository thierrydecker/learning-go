package main

import (
	"fmt"
)

func main() {

	defer func() {
		message := recover()
		fmt.Printf("Message -> %v", message)
	}()

	myString := "abcdefghijk"

	for i := 0; i <= len(myString)+1; i++ {
		fmt.Printf("Letter[%v] is %q\n", i, myString[i])
	}

}
