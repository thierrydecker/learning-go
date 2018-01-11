package main

import (
	"fmt"
)

func main() {
	myRectangle := Rectangle{width: 2, length: 5}
	fmt.Printf(
		"Area of rectangle{width: %f,length: %f} is %f\n",
		myRectangle.width, myRectangle.length, myRectangle.area())
	mySquare := Square{side: 10}
	fmt.Printf(
		"Area of square{side %f} is %f\n",
		mySquare.side, mySquare.area())
	myCircle := Circle{radius: 2}
	fmt.Printf(
		"Area of circle{radius %f} is %f\n",
		myCircle.radius, myCircle.area())
	myTriangle := Triangle{base: 10, height: 1}
	fmt.Printf(
		"Area of triangle{base: %f,height: %f} is %f\n",
		myTriangle.base, myTriangle.height, myTriangle.area())
	myShapes := []Shape{myRectangle, myTriangle, mySquare, myCircle}
	fmt.Printf("Total area is %f\n", totalArea(myShapes))
}
