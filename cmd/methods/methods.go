package main

import (
	"math"
	"fmt"
)

type Rectangle struct {
	width  float64
	length float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.length * r.width
}

func (r Circle) area() float64 {
	return math.Pi * r.radius * r.radius
}

func main() {
	var myRectangle Rectangle
	myRectangle.width = 10
	myRectangle.length = 100
	fmt.Printf(
		"Area of Rectangle(%v,%v) is %v\n",
		myRectangle.width, myRectangle.length, myRectangle.area())

	myCircle := new(Circle)
	myCircle.radius = 10
	fmt.Printf(
		"Area of Circle(%v) is %v\n",
		myCircle.radius, myCircle.area())
}
