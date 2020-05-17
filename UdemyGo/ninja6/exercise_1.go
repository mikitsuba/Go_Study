package main

import (
	"fmt"
)

func main() {
	x := foo()
	fmt.Println(x)
	y, z := bar()
	fmt.Println(y)
	fmt.Println(z)
}

func foo() int {
	return 32
}

func bar() (int, string) {
	x := 3
	y := "James"
	return x, y
}
