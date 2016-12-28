package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

// Here we use a channel to interchange information and a WaitGroup to keep the
// main thread alive until sub threads finish.
var channelA = make(chan int, 5) // uncomment this to test a buffered channel
//var channelA = make(chan int) // unbuffered channel, same as using: make(chan int, 0)
var wg sync.WaitGroup // Our friend for multi threads

// Random number generator
var ownRand *rand.Rand

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // GOMAXPROC = CPU qty.
}

func main() {
	rs := rand.NewSource(time.Now().UnixNano())
	ownRand = rand.New(rs)

	wg.Add(2)            // consider two threads
	go writeToChannel()  // one for write into the channel
	go readFromChannel() // another one for read from the channel
	wg.Wait()            // until everything goes done
}

func readFromChannel() {
	for item := range channelA { // While channel contains elements
		//fmt.Println("--> Reading")
		fmt.Print(string(item)) //print in a string format
		sleepMe(rand.Intn(100))
	}
	fmt.Print("\n")
	wg.Done() // if the channlel is closed an empty then this thread is also done
}

func writeToChannel() {
	limit := ownRand.Intn(1000) + 100 // Getting a chars limit (min 1, max 100)
	fmt.Println("Total:", limit)      // prints the limit

	for i := 0; i < limit; i++ {
		//fmt.Println("--> Writing")
		channelA <- ownRand.Intn(122-65) + 65 // Adding a random char to the channel
		sleepMe(10)
	}

	close(channelA) // Close the channel when done
	wg.Done()       // The thread is also done
}

func sleepMe(ms int) {
	sleepTime := time.Duration(ms) * time.Millisecond
	time.Sleep(sleepTime)
}
