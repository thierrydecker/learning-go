package main

import (
	"fmt"
	"github.com/thierrydecker/learning-go/pkg/math"
)

func main() {
	var a int32 = 10
	var b int32 = 10
	fmt.Printf("Sum of %d and %d is: %d\n", a, b, math.Add(a, b))
}
