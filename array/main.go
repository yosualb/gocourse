package main

import "fmt"

func main() {
	// Initialized array's value
	v := [3]int{1, 2, 3}

	fmt.Println("v:", v, "len(v):", len(v), "cap(v):", cap(v))

	// Zero value of array with selected length
	x := new([3]int)

	// Set array's value
	x[0] = 10

	fmt.Println("x:", x, "len(x):", len(x), "cap(x):", cap(x))

	// Access array's value
	fmt.Println("x[0]:", x[0])

	// Range keyword
	for key, val := range x {
		fmt.Println("key:", key, "val:", val)
	}

	// Nested array
	y := [3][3]int{
		[3]int{1, 2, 3},
		[3]int{4, 5, 6},
		[3]int{7, 8, 9},
	}

	fmt.Println("y:", y, "len(y):", len(y), "cap(y):", cap(y))
}
