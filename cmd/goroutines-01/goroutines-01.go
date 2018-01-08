package main

import (
	"fmt"
	"time"
	"math/rand"
)

func f(n int) {
	for i := 0; i < 3; i++ {
		fmt.Printf("n: %v, i: %v\n", n, i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func main() {
	fmt.Printf("Press enter to stop...\n")
	for n := 0; n < 3; n++ {
		go f(n)
	}
	input := ""
	fmt.Scanln(&input)
}
