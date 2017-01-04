package main

import (
	"fmt"
	"sync"
)

// Wait group
var wg sync.WaitGroup

func foo() {
	for i := 0; i < 1000; i++ {
		fmt.Println("i:", i)
	}

	// Decrement wait group counter
	wg.Done()
}

func bar() {
	for j := 0; j < 1000; j++ {
		fmt.Println("j:", j)
	}

	// Decrement wait group counter
	wg.Done()
}

func main() {
	// Add wait group counter
	wg.Add(2)

	go foo()
	go bar()

	// Wait group until counter is zero
	wg.Wait()
}
