package main

import (
	"fmt"
	"runtime"
)

func main() {
	for i := 10; i < 20; i++ {

		switch i%2 == 0 {
		case true:
			fmt.Printf("%d id even\n", i)
		default:
			fmt.Printf("%d id odd\n", i)
		}

	}

	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.", os)
	}
	fmt.Println()

	switch {
	case runtime.GOOS == "windows":
		fmt.Print("Go runs on windows again!\n")
	}

}
