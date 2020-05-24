package main

import "fmt"

type person struct {
	first string
	last  string
}

func changeMe(p *person) {
	(*p).first = "James"
	//p.first = "James" also works.
	(*p).last = "Bond"
}

func main() {
	p1 := person{
		first: "Tsubasa",
		last:  "Miki",
	}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}
