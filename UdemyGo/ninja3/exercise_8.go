package main

import "fmt"

func main() {
	x := 4
	switch {
	case x%4 == 0:
		{
			fmt.Printf("%d can be divided by 4\n", x)
		}
	}
}
