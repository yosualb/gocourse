package main

import "fmt"

func main() {
	// Make
	a := make([]string, 3)
	m := make(map[string]int)
	c := make(chan string)

	fmt.Println("a:", a)
	fmt.Println("m:", m)
	fmt.Println("c:", c)
}
