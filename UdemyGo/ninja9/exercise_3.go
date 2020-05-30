package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var counter int

	var wg sync.WaitGroup
	const goroutines int = 100
	wg.Add(goroutines)

	for i := 1; i <= goroutines; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			fmt.Println(counter)

			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("The final count is:", counter)
}
