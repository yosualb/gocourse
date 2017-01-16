package main

import (
	"fmt"
	"time"
)

func Supply(name string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- fmt.Sprintf("Hello %v!", name)
			time.Sleep(100 * time.Millisecond)
		}
	}()
	return c
}

// Timeout for long request
func DeliverTimeout(input <-chan string) {
	go func() {
		for {
			select {
			case s := <-input:
				fmt.Println(s)
			case <-time.After(101 * time.Millisecond):
				fmt.Println("DeliverTimeout: Input timeout")
				return
			}
		}
	}()
}

// Timeout for timed request
func DurationTimeout(input <-chan string) {
	go func() {
		timeout := time.After(1 * time.Second)
		for {
			select {
			case s := <-input:
				fmt.Println(s)
			case <-timeout:
				fmt.Println("DurationTimeout: Input timeout")
				return
			}
		}
	}()
}

func main() {
	john := Supply("John")
	DeliverTimeout(john)
	james := Supply("James")
	DurationTimeout(james)

	var input string
	fmt.Scanln(&input)

	panic("print stack")
}
