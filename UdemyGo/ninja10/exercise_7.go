package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	ch := make(chan int)

	go func() {

		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {

				for n := 0; n < 10; n++ {
					ch <- n
				}

				wg.Done()
			}()
		}

		wg.Wait()

		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}
}
