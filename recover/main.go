package main

import "fmt"

func panicTest() {
	defer func() {
		// Recover
		rec := recover()
		fmt.Println(rec)
		fmt.Println("After recover")
	}()

	fmt.Println("Before panic")
	panic("Panic!")
	fmt.Println("After panic")
}

func main() {
	fmt.Println("Before panicTest()")
	panicTest()
	fmt.Println("After panicTest()")
}
