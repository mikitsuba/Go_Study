package main

import "fmt"

type person struct {
	first    string
	last     string
	iceCream []string
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		iceCream: []string{
			"chocolate",
			"banana",
		},
	}

	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		iceCream: []string{
			"vanilla",
			"strawverry",
		},
	}

	fmt.Println(p1, "\t", p2)

	fmt.Println(p1.first, p1.last)
	for i, v := range p1.iceCream {
		fmt.Println(i, "\t", v)
	}

	fmt.Println(p2.first, p2.last)
	for i, v := range p2.iceCream {
		fmt.Println(i, "\t", v)
	}
}
