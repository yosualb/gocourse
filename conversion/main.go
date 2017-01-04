package main

import "fmt"

func main() {
	var i int = 32

	// Conversion
	var j float64 = float64(i)

	fmt.Println("j:", j)

	var s string = "Hello"

	// Conversion
	var b []byte = []byte(s)

	fmt.Println("b:", b)
}
