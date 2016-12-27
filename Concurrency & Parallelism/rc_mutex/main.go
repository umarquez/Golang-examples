package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter int
var mutex sync.Mutex // MutEx = Mutual Exclusion

// For RC valitation execute using: go run -race main.go
func main() {
	wg.Add(2)

	go incrementor("Foo:") // Possible race condition
	go incrementor("Bar:") // in those twoo calls...

	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incrementor(s string) {
	for i := 0; i < 50; i++ {
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)

		// This prevents both threads writes at the same time,
		// comment mutex.Lock() and mutex.Unlock() for testing.
		mutex.Lock() // Locks while writes, holds other threads until mutex.Unlock()
		counter++
		fmt.Println(s, i, "Counter:", counter)
		mutex.Unlock() // Flushing mutex lock
	}

	wg.Done()
}
