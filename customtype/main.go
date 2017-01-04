package main

import "fmt"

// Custom type
type foo int64

func main() {
	var f foo = 64
	fmt.Println("f:", f)

	var i int64 = 32
	// Parsing to custom type
	f = foo(i)
	fmt.Println("f:", f)
}
