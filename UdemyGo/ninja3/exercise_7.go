package main

import "fmt"

func main() {
	i := 8
	if i%4 == 0 {
		fmt.Printf("%d can be divided by 4\n", i)
	} else if i%8 == 0 {
		fmt.Printf("%d can be divided by 8\n", i)
	} else {
		fmt.Printf("%d can not be divided by 4 or 8\n", i)
	}
}
