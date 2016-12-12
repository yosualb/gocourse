package main

import "fmt"

func main() {
	// Shorthand variable
	i := 1

	// Zero value variable
	var j int

	// Initialized variable
	var k int = 10

	// Multi-variable
	var (
		l int    = 10
		m string = "A"
	)

	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)
	fmt.Println(m)
}
