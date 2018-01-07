package main

import (
	"fmt"
)

type Rectangle struct {
	width  float64
	length float64
}

type Parallelogram struct {
	Rectangle
	height float64
}

func (r Rectangle) area() float64 {
	return r.width * r.length
}

func (p Parallelogram) volume() float64 {
	return p.area() * p.height
}

func main() {
	width := 1.
	length := 1.
	height := 1.
	myParallelogram := Parallelogram{Rectangle{width, length}, height}
	fmt.Printf(
		"Parallelogram(%v,%v,%v) volume is %v\n",
		width, length, height, myParallelogram.volume())
	fmt.Printf(
		"Parallelogram width is %v and length is %v\n",
		myParallelogram.Rectangle.width, myParallelogram.Rectangle.length)
}
