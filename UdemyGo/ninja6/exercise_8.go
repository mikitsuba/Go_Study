package main

import (
	"fmt"
)

func main() {
	f := foo()
	f()
}

func foo() func() {
	g := func() {
		fmt.Println("This is from inside of func foo()")
	}
	return g
}
