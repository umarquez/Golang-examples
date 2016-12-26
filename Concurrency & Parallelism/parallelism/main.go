package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup
var sleepLimit = 500

// -- Special initialicer func, runs befor running main func.
func init() {
	fmt.Println(runtime.NumCPU(), " CPUs.") // Prints number of CPUs
	runtime.GOMAXPROCS(runtime.NumCPU())    // GOMAXPROC = CPU qty.
}

func main() {
	times := 50
	// -- With parallelistm:

	wg.Add(2)     // Adding twoo threads
	go foo(times) // Run in a thread
	go bar(times) // Run in another thread
	wg.Wait()     // Wait until threads finish

	fmt.Println("That's all folks...")
}

func foo(loop int) {
	for i := 0; i <= loop; i++ {
		x := rand.Intn(sleepLimit)
		fmt.Println("Foo - ", x, "ms.")
		time.Sleep(time.Duration(x) * time.Millisecond) // Making some time...
	}

	wg.Done() // Thread done
}

func bar(loop int) {
	for i := 0; i <= loop; i++ {
		x := rand.Intn(sleepLimit)
		fmt.Println("Bar - ", x, "ms.")
		time.Sleep(time.Duration(x) * time.Millisecond) // Making some time...
	}

	wg.Done() // Thread done
}
