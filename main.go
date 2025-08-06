package main

import "fmt"

type triangle struct {
	base   float64
	height float64
}

type square struct {
	sideLength float64
}

type shape interface {
	area() float64
}

func (t triangle) area() float64 {
	return 0.5 * t.base * t.height
}

func (s square) area() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Printf("Area is: %v", s.area())
}

func main() {
	t := triangle{
		base:   12.5,
		height: 12,
	}
	s := square{
		sideLength: 5.0,
	}

	printArea(t)
	printArea(s)
}
