package main

import "fmt"

// Struct
type Person struct {
	name string
	age  int64
}

// Empty struct
type Empty struct{}

func main() {
	// Initialize struct
	p := Person{"Yosua", 22}
	fmt.Println("p:", p)

	// Initialize struct
	e := &Empty{}
	fmt.Println("e:", e)
}
