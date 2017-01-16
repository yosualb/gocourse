package main

import "fmt"
import "time"

func Player(name string, done chan bool) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case <-done:
				return
			case c <- fmt.Sprintf("%v: Tap", name):
			}
		}
	}()
	return c
}

func System(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()
	return c
}

func main() {
	// Semaphore channel
	done := make(chan bool)

	p1 := Player("Player 1", done)
	p2 := Player("Player 2", done)
	c := System(p1, p2)

	fmt.Println("Round 1:")
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
		time.Sleep(100 * time.Millisecond)
	}
	done <- true

	fmt.Println("Round 2:")
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
		time.Sleep(100 * time.Millisecond)
	}

	var input string
	fmt.Scanln(&input)
}
