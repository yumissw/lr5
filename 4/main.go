package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func Area(s Shape) {
	fmt.Printf("площадь: %.2f", s.Area())
}

func main() {
	r := Rectangle{width: 5, height: 10}
	c := Circle{radius: 7}

	fmt.Println("длина прямоугольника: ", r.height)
	fmt.Println("ширина прямоугольника: ", r.width)
	Area(r)

	fmt.Println("\nрадиус круга: ", c.radius)
	Area(c)
}
