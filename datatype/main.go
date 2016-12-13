package main

import "fmt"

func main() {
	// Generic integer
	var a int = 1

	// 8-bit integer
	var b int8 = 2

	// 16-bit integer
	var c int16 = 3

	// 32-bit integer
	var d int32 = 4

	// 64-bit integer
	var e int64 = 5

	// Unsigned generic integer
	var f uint = 6

	// Unsigned 8-bit integer
	var g uint8 = 7

	// Unsigned 16-bit integer
	var h uint16 = 8

	// Unsigned 32-bit integer
	var i uint32 = 9

	// Unsigned 64-bit integer
	var j uint64 = 10

	// 32-bit floating-point
	var k float32 = 11.5

	// 64-bit floating-point
	var l float64 = 12.5

	// 32-bit float and 32-bit imaginary part
	var m complex64 = 13

	// 64-bit float and 64-bit imaginary part
	var n complex128 = 14

	// Alias for int32 to store UTF-8 code
	var o rune = 15

	// Alias for uint8 to store byte
	var p byte = 16

	// Unsigned large integer to store uninterpreted bits of pointer value
	var q uintptr = 17

	// Boolean
	var r bool = true

	// String
	var s string = "A"

	fmt.Println("int:", a)
	fmt.Println("int8:", b)
	fmt.Println("int16:", c)
	fmt.Println("int32:", d)
	fmt.Println("int64:", e)
	fmt.Println("uint:", f)
	fmt.Println("uint8:", g)
	fmt.Println("uint16:", h)
	fmt.Println("uint32:", i)
	fmt.Println("uint64:", j)
	fmt.Println("float32:", k)
	fmt.Println("float64:", l)
	fmt.Println("complex64:", m)
	fmt.Println("complex128:", n)
	fmt.Println("rune:", o)
	fmt.Println("byte:", p)
	fmt.Println("uintptr:", q)
	fmt.Println("bool:", r)
	fmt.Println("string:", s)
}
