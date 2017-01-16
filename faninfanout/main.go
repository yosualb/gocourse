package main

import (
	"fmt"
	"strconv"
	"time"
)

func FanIn(input1, input2 <-chan int) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case i := <-input1:
				c <- fmt.Sprintf("Input1: %v", strconv.Itoa(i))
			case i := <-input2:
				c <- fmt.Sprintf("Input2: %v", strconv.Itoa(i))
			}
		}
	}()
	return c
}

func SupplyValue(is []int) <-chan int {
	c := make(chan int)
	go func() {
		for _, val := range is {
			c <- val
			time.Sleep(100 * time.Millisecond)
		}
		close(c)
	}()
	return c
}

func Double(input <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for n := range input {
			c <- n * 2
			time.Sleep(100 * time.Millisecond)
		}
	}()
	return c
}

func main() {
	is := []int{1, 2, 3, 4, 5}
	val := SupplyValue(is)

	// Fan-out pattern
	dou1 := Double(val)
	dou2 := Double(val)

	// Fan-in pattern
	c := FanIn(dou1, dou2)

	for i := 0; i < len(is); i++ {
		fmt.Println(<-c)
	}

	var input string
	fmt.Scanln(&input)
}
