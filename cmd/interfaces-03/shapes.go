package main

import (
	"math"
)

type Rectangle struct {
	width  float64
	length float64
}

type Square struct {
	side float64
}

type Triangle struct {
	base   float64
	height float64
}

type Circle struct {
	radius float64
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

func (t Triangle) area() float64 {
	return (t.base * t.height) / 2
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func totalArea(s []Shape) float64 {
	total := 0.
	for _, shape := range s {
		total += shape.area()
	}
	return total
}
