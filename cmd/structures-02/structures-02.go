package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x float64
	y float64
	r float64
}

type Rectangle struct {
	x float64
	y float64
	l float64
	w float64
}

func circleArea(c Circle) float64 {
	return math.Pi * c.r * c.r
}

func rectangleArea(r Rectangle) float64 {
	return r.l * r.w
}

func main() {
	c := Circle{x: 0, y: 0, r: 1}
	fmt.Printf(
		"Circle area (x: %v, y: %v, r: %v) is %v\n",
		c.r, c.y, c.r, circleArea(c))
	r := Rectangle{0, 0, 2, 10}
	r.x = 0
	r.y = 0
	r.l = 2
	r.w = 2
	fmt.Printf(
		"Rectangle area (x: %v, y: %v, l: %v, w: %v) is %v\n",
		r.x, r.y, r.l, r.w, rectangleArea(r))
}
