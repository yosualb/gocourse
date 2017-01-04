package main

import "fmt"

type Person struct {
	Name string
	Age  int64
}

// Pointer receiver method
func (p *Person) ChangeName(name string) {
	p.Name = name
}

// Value receiver method
func (p Person) SayHi(name string) {
	fmt.Println("Hi", name)
}

func main() {
	p := &Person{
		"Yosua",
		22,
	}
	// Call method
	p.ChangeName("John")
	fmt.Println("p:", p)
	// Call method
	p.SayHi("James")
}
