// MIT License

// Copyright (c) 2016 Yosua Lijanto Binar

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package main

import "fmt"

func main() {
	// Decimal
	fmt.Println(65)

	// Binary
	fmt.Printf("%b\n", 65)

	// Octa-decimal
	fmt.Printf("%o\n", 65)

	// Hexa-decimal
	fmt.Printf("%x\n", 65)

	// UTF-8 symbol
	fmt.Printf("%U\n", 65)

	// UTF-8 character
	fmt.Printf("%c\n", 65)

	// UTF-8 character (escaped)
	fmt.Printf("%q\n", 65)

	// UTF-8 code
	fmt.Println('A')
}
