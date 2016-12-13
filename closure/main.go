package main

import "fmt"

// Function-level closure
func number(x int) func() int {
	return func() int {
		x++
		return x
	}
}

func main() {
	add := number(1)
	fmt.Println("add():", add())
	fmt.Println("add():", add())

	// Anonymous-function-level closure
	number2 := func(x int) func() int {
		return func() int {
			x++
			return x
		}
	}

	add2 := number2(1)
	fmt.Println("add2():", add2())
	fmt.Println("add2():", add2())
}
