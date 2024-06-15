package main

import "fmt"

type Circle struct {
	radius float64
}

var pi = 3.1459

func (c Circle) Circumference() float64 {
	return 2 * pi * c.radius
}

func (c Circle) Area() float64 {
	return pi * c.radius * c.radius
}

func main() {

	example := Circle{radius: 4}

	fmt.Sprintf("The circumference of the circle is: %f", example.Circumference())
	fmt.Sprintf("The area of the circle is: %f", example.Area())

}
