package main

import "fmt"

func main() {
	// Error, variable err is not used
	// n, err := fmt.Println("Hello World")

	// Blank identifier
	n, _ := fmt.Println("Hello World")
	fmt.Println("n:", n)
}
