package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var counter int64

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		// Add variable atomicly
		atomic.AddInt64(&counter, 1)
		// Access variable atomicly
		fmt.Println(s, "Loop:", i, "Counter:", atomic.LoadInt64(&counter))
		time.Sleep(1 * time.Microsecond)
	}
}

func main() {
	go incrementor("Foo:")
	go incrementor("Bar:")

	var input string
	fmt.Scanln(&input)
}
