package main

import "fmt"

func main() {
	// Decimal
	fmt.Println("decimal:", 65)

	// Binary
	fmt.Printf("binary: %b\n", 65)

	// Octa-decimal
	fmt.Printf("octa-decimal: %o\n", 65)

	// Hexa-decimal
	fmt.Printf("hexa-decimal: %x\n", 65)

	// UTF-8 symbol
	fmt.Printf("utf-8 symbol: %U\n", 65)

	// UTF-8 character
	fmt.Printf("utf-8 character: %c\n", 65)

	// UTF-8 character (escaped)
	fmt.Printf("utf-8 character (escaped): %q\n", 65)

	// UTF-8 code
	fmt.Println("utf-8 code:", 'A')
}
