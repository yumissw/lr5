package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func (p Person) info() string {
	return fmt.Sprintf("имя: %s \nвозраст: %d", p.name, p.age)
}

func main() {
	p := Person{name: "Даша", age: 19}
	fmt.Println(p.info())
}
