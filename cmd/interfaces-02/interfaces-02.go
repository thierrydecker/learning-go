package main

import (
	"fmt"
)

type Rectangle struct {
	width  float64
	length float64
}

type Square struct {
	side float64
}

type Shape interface {
	area() float64
}

func (r Rectangle) area() float64 {
	return r.width * r.length
}

func (s Square) area() float64 {
	return s.side * s.side
}

func totalArea(s []Shape) float64 {
	total := 0.
	for _, shape := range s {
		total += shape.area()
	}
	return total
}

func main() {
	myRectangle := Rectangle{width: 2, length: 5}
	fmt.Printf(
		"Area of rectangle{width: %f,length: %f} is %f\n",
		myRectangle.width, myRectangle.length, myRectangle.area())
	mySquare := Square{side: 10}
	fmt.Printf(
		"Area of square{side %f} is %f\n",
		mySquare.side, mySquare.area())
	myShapes := []Shape{myRectangle, mySquare}
	fmt.Printf("Total area is %f\n", totalArea(myShapes))
}
