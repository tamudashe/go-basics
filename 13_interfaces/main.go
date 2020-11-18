package main

import (
	"fmt"
	"math"
)

// Define interface
type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	length, width float64
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (r Rectangle) area() float64 {
	return r.length * r.width
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	test1 := Circle{
		x:      10,
		y:      10,
		radius: 7,
	}

	test2 := Rectangle{
		length: 12,
		width:  7,
	}

	fmt.Printf("Cicle area: %f\n", getArea(test1))
	fmt.Printf("Rectangle area: %f\n", getArea(test2))
}
