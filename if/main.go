package main

import "fmt"

func main() {
	i := 0
	// If statement
	if i < 10 {
		fmt.Println("True")
	}

	// If-else statement
	if i > 0 {
		fmt.Println("i is greater than 0")
	} else {
		fmt.Println("i is less than or equal 0")
	}

	// If-elseif-else statement
	if i > 0 {
		fmt.Println("i is greater than 0")
	} else if i < 0 {
		fmt.Println("i is less than 0")
	} else {
		fmt.Println("i is equal 0")
	}

	// Variable declaration
	if _, err := fmt.Println("Variable declaration"); err != nil {
		fmt.Println(err)
	}
}
