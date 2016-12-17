package main

import "fmt"

func main() {
	// Arithmetic operator
	// - Sum
	fmt.Println("1+2:", 1+2)

	// - Difference
	fmt.Println("1-2:", 1-2)

	// - Product
	fmt.Println("1*2:", 1*2)

	// - Quotient
	fmt.Println("1/2:", 1/2)

	// - Remainder
	fmt.Println("1%2:", 1%2)

	// - Bitwise AND
	fmt.Println("1&2:", 1&2)

	// - Bitwise OR
	fmt.Println("1|2:", 1|2)

	// - Bitwise XOR
	fmt.Println("1^2:", 1^2)

	// - Bit clear AND NOT
	fmt.Println("1&^2:", 1&^2)

	// - Left shift
	fmt.Println("1<<2:", 1<<2)

	// - Right shift
	fmt.Println("1>>2:", 1>>2)

	// Unary operator
	// - Positive number
	fmt.Println("+1:", +1)

	// - Negation
	fmt.Println("-1", -1)

	// - Bitwise complement
	fmt.Println("^1:", ^1)

	// Comparison operator
	// - Equal
	fmt.Println("1 == 2:", 1 == 2)

	// - Not equal
	fmt.Println("1 != 2:", 1 != 2)

	// - Less
	fmt.Println("1 < 2:", 1 < 2)

	// - Less or equal
	fmt.Println("1 <= 2:", 1 <= 2)

	// - Greater
	fmt.Println("1 > 2:", 1 > 2)

	// - Greater or equal
	fmt.Println("1 >= 2:", 1 >= 2)

	// Logical operator
	// - AND
	fmt.Println("true && false:", true && false)

	// - OR
	fmt.Println("true || false:", true || false)

	// - NOT
	fmt.Println("!true:", !true)

	// Address operator
	x := 10
	y := &x
	fmt.Println("x:", x)
	fmt.Println("&x:", &x)
	fmt.Println("y:", y)
	fmt.Println("&y:", &y)
	fmt.Println("*y:", *y)

	// Receive operator
	c := make(chan int, 1)
	c <- 1
	close(c)
	fmt.Println("c:", <-c)

	// Assignment operator
	i := 1

	// - Addition
	i += 2
	fmt.Println("i:", i)

	// - Substraction
	i -= 2
	fmt.Println("i:", i)

	// - Multiplication
	i *= 2
	fmt.Println("i:", i)

	// - Division
	i /= 2
	fmt.Println("i:", i)

	// - Remainder
	i %= 2
	fmt.Println("i:", i)

	// - Bitwise AND
	i &= 2
	fmt.Println("i:", i)

	// - Bitwise OR
	i |= 2
	fmt.Println("i:", i)

	// - Bitwise XOR
	i ^= 2
	fmt.Println("i:", i)

	// - Left shift
	i <<= 2
	fmt.Println("i:", i)

	// - Right shift
	i >>= 2
	fmt.Println("i:", i)

	// - Bit clear AND NOT
	i &^= 2
	fmt.Println("i:", i)

	// IncDec operator
	j := 1

	// - Increment
	j++
	fmt.Println("j:", j)

	// - Decrement
	j--
	fmt.Println("j:", j)
}
