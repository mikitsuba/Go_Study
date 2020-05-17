package main

import (
	"fmt"
)

func main() {
	defer foo()
	bar()
}

func foo() {
	fmt.Println("This is from foo.")
}

func bar() {
	fmt.Println("This is from bar.")
}
