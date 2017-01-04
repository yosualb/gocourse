package main

import (
	"fmt"
	"reflect"
)

func main() {
	var s interface{} = "Hello"

	// Assertion
	if v, ok := s.(string); ok {
		fmt.Println("v:", reflect.TypeOf(v))
	} else {
		fmt.Println("Not string")
	}
}
