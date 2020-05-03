package main

import (
	"fmt"
)

func main() {
	p := struct {
		first string
		last  string
		age   int
	}{
		first: "Tsubasa",
		last:  "Miki",
		age:   30,
	}
	fmt.Println(p)
	fmt.Printf("%v%T\n", p, p)
}
