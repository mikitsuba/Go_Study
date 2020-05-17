package main

import (
	"fmt"
)

func main() {
	f := foo
	f()
}

func foo() {
	fmt.Println("This is foo")
}
