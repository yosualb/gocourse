package main

import (
	"fmt"
	"time"
)

// Generator pattern
func Generator(name string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- fmt.Sprintf("Hello %v!", name)
			time.Sleep(100 * time.Millisecond)
		}
	}()
	return c
}

func main() {
	john := Generator("John")

	go func() {
		for n := range john {
			fmt.Println("n:", n)
		}
	}()

	var input string
	fmt.Scanln(&input)
}
