package main

import (
	"fmt"
	"sync"
	"time"
)

var counter int64

// Mutex
var mutex sync.Mutex

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		mutex.Lock()
		counter++
		time.Sleep(1 * time.Microsecond)
		fmt.Println(s, "Loop:", i, "Counter:", counter)
		mutex.Unlock()
	}
}

func main() {
	go incrementor("Foo:")
	go incrementor("Bar:")

	var input string
	fmt.Scanln(&input)
}
