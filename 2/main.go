package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func (p *Person) birthday() {
	p.age++
}

func main() {

	p := Person{name: "Даша", age: 19}
	fmt.Println("до: ", p.age)
	p.birthday()
	fmt.Println("после: ", p.age)
}
