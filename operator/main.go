package main

import "fmt"

func main() {
	// Arithmetic operator
	// - Sum
	fmt.Println("1+1:", 1+1)

	// - Difference
	fmt.Println("1-1:", 1-1)

	// - Product
	fmt.Println("1*1:", 1*1)

	// - Quotient
	fmt.Println("1/1:", 1/1)

	// - Remainder
	fmt.Println("1%1:", 1%1)

	// - Bitwise AND
	fmt.Println("1&1:", 1&1)

	// - Bitwise OR
	fmt.Println("1|1:", 1|1)

	// - Bitwise XOR
	fmt.Println("1^1:", 1^1)

	// - Bit clear AND NOT
	fmt.Println("1&^1:", 1&^1)

	// - Left shift
	fmt.Println("1<<1:", 1<<1)

	// - Right shift
	fmt.Println("1>>1:", 1>>1)

	// Unary operator
	// - Positive number
	fmt.Println("+1:", +1)

	// - Negation
	fmt.Println("-1", -1)

	// - Bitwise complement
	fmt.Println("^1:", ^1)

	// Comparison operator
	// - Equal
	fmt.Println("1 == 1:", 1 == 1)

	// - Not equal
	fmt.Println("1 != 1:", 1 != 1)

	// - Less
	fmt.Println("1 < 1:", 1 < 1)

	// - Less or equal
	fmt.Println("1 <= 1:", 1 <= 1)

	// - Greater
	fmt.Println("1 > 1:", 1 > 1)

	// - Greater or equal
	fmt.Println("1 >= 1:", 1 >= 1)

	// Logical operator
	// - AND
	fmt.Println("true && false:", true && false)

	// - OR
	fmt.Println("true || false:", true || false)

	// - NOT
	fmt.Println("!true:", !true)

	// Address operator
	var x int = 10
	var y *int = &x
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
	var i int = 1

	// - Addition
	i += 1
	fmt.Println("i:", i)

	// - Substraction
	i -= 1
	fmt.Println("i:", i)

	// - Multiplication
	i *= 1
	fmt.Println("i:", i)

	// - Division
	i /= 1
	fmt.Println("i:", i)

	// - Remainder
	i %= 1
	fmt.Println("i:", i)

	// - Bitwise AND
	i &= 1
	fmt.Println("i:", i)

	// - Bitwise OR
	i |= 1
	fmt.Println("i:", i)

	// - Bitwise XOR
	i ^= 1
	fmt.Println("i:", i)

	// - Left shift
	i <<= 1
	fmt.Println("i:", i)

	// - Right shift
	i >>= 1
	fmt.Println("i:", i)

	// - Bit clear AND NOT
	i &^= 1
	fmt.Println("i:", i)

	// IncDec operator
	var j int = 1

	// - Increment
	j++
	fmt.Println("j:", j)

	// - Decrement
	j--
	fmt.Println("j:", j)
}
