package main

import "fmt"

// Function without receiver, argument, and return value
func hi() {
	fmt.Println("Hello")
}

// Function with return value
func greet() string {
	return "Hello World"
}

// Function with argument and return value
func greetPlus(s string) string {
	return "Hello " + s
}

// Variadic function
func talk(ss ...string) string {
	var names string
	for key, val := range ss {
		if key == len(ss)-1 {
			names += val
			continue
		}
		names += val + ", "
	}
	return "Hi " + names
}

// Function with named return value
func greeting() (greeting string) {
	greeting = "Hello World"
	return
}

func main() {
	hi()
	fmt.Println("greet():", greet())
	fmt.Println("greetPlus(\"Yosua\"):", greetPlus("Yosua"))
	fmt.Println("talk(\"Yosua\", \"James\", \"Ben\"):", talk("Yosua", "James", "Ben"))

	names := []string{"Yosua", "James", "Ben"}
	// Variadic argument
	fmt.Println("talk(names):", talk(names...))

	// Function expression
	hello := func() string {
		return "Hello World"
	}

	fmt.Println("hello():", hello())

	// Anonymous self-executing function
	func() {
		fmt.Println("Hello World")
	}()

	fmt.Println("greeting():", greeting())
}
