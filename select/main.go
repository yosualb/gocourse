package main

import "fmt"

func Player(name string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < 3; i++ {
			c <- fmt.Sprintf("%v is tapping", name)
		}
		close(c)
	}()
	return c
}

func GameSystem(input1, input2 <-chan string) {
	go func() {
		for {
			if input1 == nil && input2 == nil {
				break
			}
			// Select statement
			select {
			case s, ok := <-input1:
				if !ok {
					input1 = nil
					break
				}
				fmt.Println("input1:", s)
			case s, ok := <-input2:
				if !ok {
					input2 = nil
					break
				}
				fmt.Println("input2:", s)
			}
		}
	}()
}

func main() {
	GameSystem(Player("A"), Player("B"))

	var input string
	fmt.Scanln(&input)
}
