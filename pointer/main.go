package main

import "fmt"

// Pass memory address
func zero(x *int) {
	*x = 0
}

func main() {
	i := 10

	// Variable memory address
	j := &i

	// Change memory address's value
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
