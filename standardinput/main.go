package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Standard input
	var s string
	fmt.Scanln(&s)
	fmt.Println("s:", s)

	// Buffered standard input
	sc := bufio.NewScanner(os.Stdin)
	if sc.Scan() {
		fmt.Println(sc.Text())
	}
}
