package main

import "fmt"

func main() {
	// Print into standard output
	n, err := fmt.Println("Hello World!")
	fmt.Println("n:", n, "err:", err)
}
