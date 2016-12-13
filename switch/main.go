package main

import "fmt"

func main() {
	name := "James"
	// Switch statement
	switch name {
	case "James":
		fmt.Println("Hi I'm James")
	case "Jenny":
		fmt.Println("Hello I'm Jenny")
	default:
		fmt.Println("Who am I?")
	}

	// Fallthrough keyword executes exactly next case
	switch name {
	case "James":
		fmt.Println("Hi I'm James")
		fallthrough
	case "Jenny":
		fmt.Println("Hello I'm Jenny")
	default:
		fmt.Println("Who am I?")
	}

	// Switch statement without condition
	switch {
	case name == "James":
		fmt.Println("Hi I'm James")
	case name == "Jenny":
		fmt.Println("Hello I'm Jenny")
	default:
		fmt.Println("Who am I?")
	}

	var test interface{} = "Hello"
	// Type checking
	switch test.(type) {
	case string:
		fmt.Println("string")
	case int:
		fmt.Println("int")
	default:
		fmt.Println("<nil>")
	}
}
