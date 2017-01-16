package main

import "fmt"

func panicTest() {
	fmt.Println("Before panic")
	panic("Panic!")
	fmt.Println("After panic")
}

func main() {
	defer func() {
		// Recover
		rec := recover()
		fmt.Println("rec:", rec)
		fmt.Println("After recover")
	}()

	fmt.Println("Before panicTest()")
	panicTest()
	fmt.Println("After panicTest()")
}
