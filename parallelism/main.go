package main

import "fmt"
import "runtime"

func foo() {
	for i := 0; i < 1000; i++ {
		fmt.Println("i:", i)
	}
}

func bar() {
	for j := 0; j < 1000; j++ {
		fmt.Println("j:", j)
	}
}

// Support multiple core
func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	go foo()
	go bar()

	var input string
	fmt.Scanln(&input)
}
