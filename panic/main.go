package main

import "fmt"

func panicTest() {
	fmt.Println("Before panic")
	// Panic
	panic("Panic!")
	fmt.Println("After panic")
}

func main() {
	fmt.Println("Before panicTest()")
	panicTest()
	fmt.Println("After panicTest()")
}
