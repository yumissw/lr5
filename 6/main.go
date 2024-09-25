package main

import (
	"fmt"
)

type Stringer interface {
	String() string
}

type Book struct {
	Title  string
	Author string
}

func (b Book) String() string {
	return fmt.Sprintf("название: %s\nавтор: %s", b.Title, b.Author)
}

func main() {
	b := Book{"Бедная Лиза", "Н.М. Карамзин"}
	fmt.Println(b.String())
}
