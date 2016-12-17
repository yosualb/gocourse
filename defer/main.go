package main

import "fmt"

func main() {
	// Defer keyword
	defer fmt.Println("World")
	fmt.Println("Hello")
}
