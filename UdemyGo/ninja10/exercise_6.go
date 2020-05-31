package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}
		close(ch)
	}()

	for val := range ch {
		fmt.Println("Received from ch:", val)
	}
}
