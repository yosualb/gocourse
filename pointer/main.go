package main

import "fmt"

// Pass memory address to function
func zero(x *int) {
	*x = 0
}

func main() {
	i := 10

	// Pointer variable that assigned by variable i's memory address
	j := &i

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
