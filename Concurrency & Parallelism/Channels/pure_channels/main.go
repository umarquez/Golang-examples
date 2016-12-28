package main

import (
	"fmt"
	"math/rand"
	"time"
)

var status = make(chan bool)
var chars = make(chan int, 10)

func main() {
	maxChars := 100
	maxThreads := 10

	rs := rand.NewSource(time.Now().UnixNano())
	ownRand := rand.New(rs)

	// Creating threads:
	for threads := 0; threads < maxThreads; threads++ {
		// Anonymous functions as threads
		go func() {
			for i := 0; i < maxChars; i++ {
				chars <- ownRand.Intn(122-65) + 65
			}

			status <- true
		}()
	}

	// Status manager function:
	go func() {
		for i := 0; i < maxThreads; i++ {
			<-status
		}
		close(status)
		close(chars)
	}()

	// Printing chars
	for singleChar := range chars {
		fmt.Print(string(singleChar))
	}

	fmt.Print("\n")
}
