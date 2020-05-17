package main

import (
	"fmt"
)

func main() {
	n := factorial(12)
	fmt.Println(n)
}

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)

}
