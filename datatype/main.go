package main

import "fmt"
import "reflect"

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

	fmt.Println("a:", reflect.TypeOf(a))
	fmt.Println("b:", reflect.TypeOf(b))
	fmt.Println("c:", reflect.TypeOf(c))
	fmt.Println("d:", reflect.TypeOf(d))
	fmt.Println("e:", reflect.TypeOf(e))
	fmt.Println("f:", reflect.TypeOf(f))
	fmt.Println("g:", reflect.TypeOf(g))
	fmt.Println("h:", reflect.TypeOf(h))
	fmt.Println("i:", reflect.TypeOf(i))
	fmt.Println("j:", reflect.TypeOf(j))
	fmt.Println("k:", reflect.TypeOf(k))
	fmt.Println("l:", reflect.TypeOf(l))
	fmt.Println("m:", reflect.TypeOf(m))
	fmt.Println("n:", reflect.TypeOf(n))
	fmt.Println("o:", reflect.TypeOf(o))
	fmt.Println("p:", reflect.TypeOf(p))
	fmt.Println("q:", reflect.TypeOf(q))
	fmt.Println("r:", reflect.TypeOf(r))
	fmt.Println("s:", reflect.TypeOf(s))
}
