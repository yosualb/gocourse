package main

import "fmt"

func result(x int) int {
	return 10 + x
}

func main() {
	// Variable shadowing
	result := result(2)
	fmt.Println(result)

	// Error, function result is overshadowed by variable result
	// result = result(3)
}
