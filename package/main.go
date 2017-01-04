package main

import (
	"fmt"

	// Package alias
	s "github.com/yosualb/gocourse/package/string"
)

func main() {
	// Call package method
	fmt.Println("string.AddNamePrefix(\"Mr. \"):", s.AddNamePrefix("Mr. "))
	// Call package variable
	fmt.Println("string.Name:", s.Name)
}
