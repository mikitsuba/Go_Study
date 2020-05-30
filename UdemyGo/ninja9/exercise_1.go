package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Number of Goroutine (begin):", runtime.NumGoroutine())
	wg.Add(2)
	fmt.Println("This is the main goroutine.")

	go foo()
	fmt.Println("Number of Goroutine (after foo):", runtime.NumGoroutine())

	go bar()
	fmt.Println("Number of Goroutine (after bar):", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("Number of Goroutine (end):", runtime.NumGoroutine())
}

func foo() {
	fmt.Println("This is Foo!")
	wg.Done()
}

func bar() {
	fmt.Println("This is Bar!")
	wg.Done()
}
