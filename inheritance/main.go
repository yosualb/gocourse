package main

import "fmt"

type Person struct {
	age int64
}

type Student struct {
	// Inheritance
	Person
	code string
}

func (p *Person) Say() {
	fmt.Println("I'm person")
}

func (s *Student) Say() {
	fmt.Println("I'm student")
}

func main() {
	s := &Student{
		// Initialize inheritance
		Person{22},
		"ABC",
	}
	fmt.Println("s:", s)
	s.Say()
	// Call inherintance method
	s.Person.Say()
}
