package main

import (
	"fmt"
)

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("Here is the error: %v", ce.info)
}

func main() {
	c1 := customErr{
		info: "Nothing.",
	}
	foo(c1)
}

func foo(e error) {
	fmt.Println(e)
	fmt.Println(e.(customErr).info) //asertion
}
