package main

import (
	"fmt"
	"math"
)

// Define Interface
type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

// struct method
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

// main abstract class
// accept Struct argument as Shape Interface
//
// call Interface.method -> area()
// so find area() in Accepted struct's method
// and call it

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circle := Circle{x: 3, y: 2, radius: 5}
	rectangle := Rectangle{width: 22, height: 10}

	fmt.Printf("Circle area is: %f\n", getArea(circle))
	fmt.Printf("Rectangle area is: %f\n", getArea(rectangle))

	// rectangle.area()
}
