package main

import "fmt"

// Package-level scope variable
var i int = 10

func main() {
	// Function-level scope variable
	j := 11

	{
		// Inner-function-level scope variable
		k := 12
		fmt.Println("k:", k)
	}

	fmt.Println("i:", i)
	fmt.Println("j:", j)
	// Error, variable k is in inner scope
	// fmt.Println("k:", k)
}
