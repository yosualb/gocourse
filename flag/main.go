package main

import (
	"flag"
	"fmt"
)

func main() {
	// Define flag
	x := flag.Int64("x", 0, "X value")
	y := flag.Int64("y", 0, "Y value")

	// Register flag
	flag.Parse()

	fmt.Println("*x + *y:", *x+*y)
}
