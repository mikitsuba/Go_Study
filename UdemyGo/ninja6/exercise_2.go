package main

import (
	"fmt"
)

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	x := foo(xi...)
	fmt.Println(x)

	y := bar(xi)
	fmt.Println(y)
}

func foo(n ...int) int {
	sum := 0
	for _, v := range n {
		sum += v
	}
	return sum
}

func bar(xi []int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
