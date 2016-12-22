package main

import "fmt"

func main() {
	// New
	x := new(int)
	y := new(float64)

	fmt.Println("x:", x)
	fmt.Println("*x:", *x)
	fmt.Println("y:", y)
	fmt.Println("*y:", *y)
}
