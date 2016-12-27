package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var counter int64

// For RC valitation execute using: go run -race main.go
func main() {
	wg.Add(2)

	go incrementor("Foo:") // Possible race condition
	go incrementor("Bar:") // in those twoo calls...

	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)

		// Using atomic functions for preventing RC
		// NOTE: ensure to use the correct atomic function for the data type
		atomic.AddInt64(&counter, 1) // atomic equivalent for counter++
		fmt.Println(s, i, "Counter:", counter)
	}

	wg.Done()
}
