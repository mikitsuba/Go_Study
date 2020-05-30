package main

import "fmt"

type person struct {
	First string
	Last  string
}

func (a *person) speak() {
	fmt.Println(a.First, a.Last, "says hello.")
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	person := person{
		First: "Tsubasa",
		Last:  "Miki",
	}
	// saySomething(person) this doesn't run because type human doesn't include speak() as its method set, but type *human does.
	saySomething(&person)
	person.speak()

}
