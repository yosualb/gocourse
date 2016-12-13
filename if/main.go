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
		fmt.Println("A")
	} else {
		fmt.Println("B")
	}

	// If-elseif-else statement
	if i > 0 {
		fmt.Println("A")
	} else if i < 0 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

	// Variable declaration
	if _, err := fmt.Println("Variable declaration"); err != nil {
		fmt.Println(err)
	}
}
