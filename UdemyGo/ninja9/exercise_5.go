package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64

	var wg sync.WaitGroup
	const goroutines int = 100
	wg.Add(goroutines)

	for i := 1; i <= goroutines; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println("The count is:", atomic.LoadInt64(&counter))

			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("The final count is:", counter)
}
