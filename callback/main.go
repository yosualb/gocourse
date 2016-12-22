package main

import "fmt"

func formattedPrint(s string) {
	fmt.Println("Hello my name is", s)
}

// Function that contains callback function
func printName(callback func(string), ss ...string) {
	for _, val := range ss {
		callback(val)
	}
}

func main() {
	printName(formattedPrint, "Yosua", "James", "Ben")
	printName(func(s string) {
		fmt.Println("Hi", s)
	}, "Yosua", "James", "Ben")
}
