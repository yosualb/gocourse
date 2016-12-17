package main

import "fmt"

// Recursion
func addInt(i int) int {
	// Base case
	if i == 0 {
		return 0
	}
	// Recursive case
	return i + addInt(i-1)
}

func main() {
	fmt.Println("addInt(5):", addInt(5))
}
