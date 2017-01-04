package main

import "fmt"

func foo() {
	for i := 0; i < 1000; i++ {
		fmt.Println("i:", i)
	}
}

func bar() {
	for j := 0; j < 1000; j++ {
		fmt.Println("j:", j)
	}
}

func main() {
	// Concurrency
	go foo()
	// Concurrency
	go bar()

	var input string
	fmt.Scanln(&input)
}
