package main

import (
	"fmt"
	"os"
)

func main() {
	// Open file
	file, err := os.Open("file/input.txt")
	if err != nil {
		return
	}
	// Close file
	defer file.Close()

	// Check file status
	stat, err := file.Stat()
	if err != nil {
		return
	}

	// Read file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}

	fmt.Println("string(bs):", string(bs))

	// Create file
	file, err = os.Create("file/output.txt")
	if err != nil {
		return
	}
	// Close file
	defer file.Close()

	// Write file
	file.WriteString("Hello World!")
}
