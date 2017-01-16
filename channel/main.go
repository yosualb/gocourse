package main

import (
	"fmt"
	"time"
)

// Two-way channel
func worker(input chan string) {
	go func() {
		for i := 65; i < 75; i++ {
			input <- string(i)
			time.Sleep(100 * time.Millisecond)
		}
		// Close channel
		close(input)
	}()
}

func workerPrinter(input <-chan string) {
	go func() {
		// Range keyword
		for n := range input {
			fmt.Println("n:", n)
		}
	}()
}

// Write-only channel
func supply(input chan<- string) {
	go func() {
		for {
			input <- "Hello"
			time.Sleep(100 * time.Millisecond)
		}
	}()
}

// Read-only channel
func print(input <-chan string) {
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("<-input:", <-input)
		}
	}()
}

func main() {
	// Two-way channel
	c := make(chan string)
	c1 := make(chan string)

	worker(c)
	workerPrinter(c)

	supply(c1)
	print(c1)

	// Buffered channel
	bc := make(chan string, 2)

	go func() {
		bc <- "A"
		bc <- "B"
		bc <- "C"
		bc <- "D"
		bc <- "E"
	}()

	go func() {
		for {
			if len(bc) > 1 {
				fmt.Println("bc:", <-bc)
				time.Sleep(1 * time.Second)
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}
