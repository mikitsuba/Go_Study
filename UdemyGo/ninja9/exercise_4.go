package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int

	var wg sync.WaitGroup
	const goroutines int = 100
	wg.Add(goroutines)

	var mu sync.Mutex

	for i := 1; i <= goroutines; i++ {
		go func() {
			mu.Lock()

			v := counter
			// runtime.Gosched() it needs to be removed
			v++
			counter = v
			fmt.Println(counter)

			mu.Unlock()

			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("The final count is:", counter)
}
