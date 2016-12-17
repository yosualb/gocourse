package main

import (
	"fmt"

	"github.com/yosualb/gocourse/package/string"
)

func main() {
	// Call package's exported function
	fmt.Println("string.AddNamePrefix(\"Mr. \"):", string.AddNamePrefix("Mr. "))

	// Call package's exported variable
	fmt.Println("string.Name:", string.Name)
}
