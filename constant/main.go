package main

import "fmt"

func main() {
	// Typed constant
	const name string = "Yosua"

	// Untyped constant
	const pi = 3.14

	// Multiple constant
	const (
		age     int    = 1
		country string = "Indonesia"
	)

	// Iota keyword enumerates constant
	const (
		zero int = iota
		one
		two
		three
	)

	fmt.Println("name:", name)
	fmt.Println("pi:", pi)
	fmt.Println("age:", age)
	fmt.Println("country:", country)
	fmt.Println("zero:", zero)
	fmt.Println("one:", one)
	fmt.Println("two:", two)
	fmt.Println("three:", three)
}
