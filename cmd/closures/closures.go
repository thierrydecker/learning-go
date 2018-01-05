package main

import (
	"fmt"
)

func makeEvenGenerator() func() int {
	i := 0
	return func() (ret int) {
		ret = i
		i += 2
		return
	}
}

func main() {
	nextEven := makeEvenGenerator()
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d - %d\n", i, nextEven())
	}
}
