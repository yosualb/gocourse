package main

import "fmt"

func main() {
	// Iteration like a C for that have init, condition, and post
	for i := 0; i < 10; i++ {
		fmt.Println("i:", i)
	}

	// Iteration like a C while that only have condition
	i := 0
	for i < 10 {
		fmt.Println("i:", i)
		i++
	}

	// Nested iteration
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println("i:", i, "j:", j)
		}
	}

	// Break keyword
	for i := 0; i < 10; i++ {
		if i > 7 {
			break
		}
		fmt.Println("i:", i)
	}

	// Continue keyword
	for i := 0; i < 10; i++ {
		if i%2 == 1 {
			continue
		}
		fmt.Println("i:", i)
	}

	var r []rune = []rune("Hello World")
	// Range keyword can create loop for array, slice, string, or map, or reading from a channel
	for key, val := range r {
		fmt.Println("key:", key, "val:", val)
	}

	// Iteration like a C for(;;) or infinite loop
	// for {
	// 	fmt.Println("Infinite loop")
	// }
}
