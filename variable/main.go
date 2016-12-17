package main

import "fmt"

func main() {
	// Shorthand variable
	i := 1

	// Zero value variable
	var j int

	// Initialized variable
	var k int = 10

	// Multiple variable
	var (
		l int64   = 10
		m float64 = 12.5
	)

	fmt.Println("i:", i)
	fmt.Println("j:", j)
	fmt.Println("k:", k)
	fmt.Println("l:", l)
	fmt.Println("m:", m)
}
