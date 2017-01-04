package main

import (
	"container/list"
	"fmt"
)

func main() {
	// List
	var l list.List

	// Add element to the last
	l.PushBack(1)
	l.PushBack(3)
	// Add element after selected element
	l.InsertAfter(2, l.Front())

	// Print all list
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println("e:", e.Value.(int))
	}
}
