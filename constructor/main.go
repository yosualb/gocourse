package main

import (
	"fmt"
	"math/rand"

	"github.com/yosualb/gocourse/constructor/lib"
)

type Person struct {
	name   string
	serial int
}

// Constructor
func NewPerson(name string) *Person {
	p := &Person{name, rand.Int()}
	return p
}

func main() {
	p := lib.NewPerson("John")
	fmt.Println("p:", p)
}
