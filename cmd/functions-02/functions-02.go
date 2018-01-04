package main

import (
	"fmt"
)

func f1() {
	fmt.Printf("Entering f1()\n")
	fmt.Printf("Executing f1()\n")
	fmt.Printf("Calling f2()\n")
	f2()
	fmt.Printf("Exiting f1()\n")
}

func f2() {
	fmt.Printf("Entering f2()\n")
	fmt.Printf("Executing f2()\n")
	fmt.Printf("Exiting f2()\n")
}

func main() {
	fmt.Printf("Entering main()\n")
	fmt.Printf("Calling f1()\n")
	f1()
	fmt.Printf("Exiting main()\n")
}
