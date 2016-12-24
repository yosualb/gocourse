package main

import "fmt"

func main() {
	// Nil map
	var a map[string]int

	fmt.Println("a == nil:", a == nil)

	// Initialized map's value
	b := map[string]int{
		"zero": 0,
	}

	fmt.Println("b:", b)

	// Zero value of map
	c := make(map[string]int)

	fmt.Println("c:", c)

	// Zero value of map with selected length of hint
	d := make(map[string]int, 10)

	fmt.Println("d:", d)

	// Set map's key and value
	d["one"] = 1

	fmt.Println("d:", d)

	// Delete map's key and value
	delete(d, "one")

	fmt.Println("d:", d)

	// Access map's key and value with availability checking
	val, ok := d["one"]
	val2, ok2 := d["two"]

	fmt.Println("val:", val, "ok:", ok)
	fmt.Println("val2:", val2, "ok2:", ok2)

	// Nested map
	e := map[string]map[string]int{
		"Tom": map[string]int{
			"money": 1000,
			"age":   22,
		},
		"Ben": map[string]int{
			"money": 2000,
			"age":   23,
		},
	}

	fmt.Println("e:", e)

	// Range keyword
	for key, val := range e {
		fmt.Println("key:", key, "val:", val)
	}
}
