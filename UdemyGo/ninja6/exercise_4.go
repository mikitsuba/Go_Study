package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "Tsubasa",
		last:  "Miki",
		age:   31,
	}

	p1.speak()

}

func (p person) speak() {
	fmt.Println("I am ", p.first, p.last, ". My age is ", p.age, ".")
}
