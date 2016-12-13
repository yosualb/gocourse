package main

import "fmt"

// Pass memory address to function
func zero(x *int) {
	*x = 0
}

func main() {
	var i int = 10

	// Pointer variable that assigned variable i's memory address
	var j *int = &i

	// Reference to the assigned memory address's value and change it
	*j = 11

	fmt.Println("i:", i)
	fmt.Println("&i:", &i)
	fmt.Println("j:", j)
	fmt.Println("*j:", *j)
	fmt.Println("&j:", &j)

	zero(j)

	fmt.Println("i:", i)
	fmt.Println("&i:", &i)
	fmt.Println("j:", j)
	fmt.Println("*j:", *j)
	fmt.Println("&j:", &j)
}
