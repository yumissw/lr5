package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {

	c := Circle{radius: 7}

	fmt.Println("радиус: ", c.radius)
	fmt.Printf("площадь: %.2f", c.Area())
}
