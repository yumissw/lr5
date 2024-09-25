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

func printAreas(shapes []Shape) {
	for _, shape := range shapes {
		switch s := shape.(type) {
		case Rectangle:
			fmt.Println("\nдлина прямоугольника: ", s.height)
			fmt.Println("ширина прямоугольника: ", s.width)
			fmt.Printf("площадь прямоугольника: %.2f\n", s.Area())
		case Circle:
			fmt.Println("\nрадиус круга: ", s.radius)
			fmt.Printf("площадь круга: %.2f\n", s.Area())
		}
	}
}

func main() {

	shapes := []Shape{
		Circle{radius: 5},
		Rectangle{width: 4, height: 6},
	}

	printAreas(shapes)

}
