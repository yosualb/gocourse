package main

import "fmt"

func main() {
	// Initialized slice's value
	v := []int{1, 2, 3}

	fmt.Println("v:", v, "len(v):", len(v), "cap(v):", cap(v))

	// Zero value of slice with selected same length and capacity
	x := make([]int, 3)

	fmt.Println("x:", x, "len(x):", len(x), "cap(x):", cap(x))

	// Zero value of slice with selected length and capacity
	y := make([]int, 3, 5)

	fmt.Println("y:", y, "len(y):", len(y), "cap(y):", cap(y))

	// Set slice's value
	y[0] = 10

	fmt.Println("y:", y, "len(y):", len(y), "cap(y):", cap(y))

	// Access slice's value
	fmt.Println("y[0]:", y[0])

	// Range keyword
	for key, val := range y {
		fmt.Println("key:", key, "val:", val)
	}

	arr := [5]int{1, 2, 3, 4, 5}
	// Create slice from slicing array with selected first and last index
	z := arr[0:3]

	fmt.Println("z:", z, "len(z):", len(z), "cap(z):", cap(z))

	// Create slice from slicing array with selected first and until last index
	a := arr[3:]

	fmt.Println("a:", a, "len(a):", len(a), "cap(a):", cap(a))

	// Create slice from slicing array from first until selected last index
	b := arr[:3]

	fmt.Println("b:", b, "len(b):", len(b), "cap(b):", cap(b))

	// Append slice
	c := append(b, 4, 5)

	fmt.Println("c:", c, "len(c):", len(c), "cap(c):", cap(c))

	d := make([]int, 2)
	// Copy slice
	copy(d, c)

	fmt.Println("d:", d, "len(d):", len(d), "cap(d):", cap(d))

	// Nested slice
	e := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}

	fmt.Println("e:", e, "len(e):", len(e), "cap(e):", cap(e))
}
