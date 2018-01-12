package main

import "fmt"

func main() {

	a := []int16{1, 2}
	b := []int16{10, 20, 30}
	c := []int16{100, 200, 300, 400}

	d := make([]int16, 0)

	d = append(d, 0)
	fmt.Printf("d: %v, len(d): %d, cap(d): %d\n", d, len(d), cap(d))

	d = append(d, a...)
	fmt.Printf("d: %v, len(d): %d, cap(d): %d\n", d, len(d), cap(d))

	d = append(d, b...)
	fmt.Printf("d: %v, len(d): %d, cap(d): %d\n", d, len(d), cap(d))

	d = append(d, c...)
	fmt.Printf("d: %v, len(d): %d, cap(d): %d\n", d, len(d), cap(d))
}
