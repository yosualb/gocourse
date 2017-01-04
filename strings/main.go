package main

import (
	"fmt"
	"strings"
)

func main() {
	// String contains
	fmt.Println("strings.Contains(\"Hello\", \"Hell\"):", strings.Contains("Hello", "Hell"))

	// String compare
	fmt.Println("strings.Compare(\"Hello\", \"HeIIo\"):", strings.Compare("Hello", "HeIIo"))

	// String count
	fmt.Println("strings.Count(\"Hello\", \"l\"):", strings.Count("Hello", "l"))

	// String has prefix
	fmt.Println("strings.HasPrefix(\"Hello\", \"Hell\"):", strings.HasPrefix("Hello", "Hell"))

	// String has suffix
	fmt.Println("strings.HasSuffix(\"Hello\", \"lo\"):", strings.HasSuffix("Hello", "lo"))

	// String index
	fmt.Println("strings.Index(\"Hello\", \"lo\"):", strings.Index("Hello", "lo"))

	// String join
	fmt.Println("strings.Join([]string{\"1\", \"2\", \"3\"}, \"-\"):", strings.Join([]string{"1", "2", "3"}, "-"))

	// String repeat
	fmt.Println("strings.Repeat(\"x\", 5):", strings.Repeat("x", 5))

	// String replace
	fmt.Println("strings.Replace(\"Hello\", \"l\", \"y\", 2):", strings.Replace("Hello", "l", "y", 2))

	// String split
	fmt.Println("strings.Split(\"a-b-c\", \"-\"):", strings.Split("a-b-c", "-"))

	// String to upper
	fmt.Println("strings.ToUpper(\"hello\"):", strings.ToUpper("hello"))

	// String to lower
	fmt.Println("strings.ToLower(\"HELLO\"):", strings.ToLower("HELLO"))
}
