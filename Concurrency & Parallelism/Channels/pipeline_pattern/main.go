package main

import "fmt"

func main() {
	numCounter := 55 // Numbers to generate

	fmt.Println("Sum of the first", numCounter, "Fibonacci numbers.")

	/**
		 * Pipeline comunication between two process:
		 *
		 * fibNumbers - Creates a thread thats generates Fibonacci numbers and stores
	   * those numbers into a channel that returns inmediatly.
	   *
	   * sumNumbers - Recive a channel and sum every number in there using a thread,
	   * the result is stored into a channel and this is returned inmediatly.
	   *
	   * Finally, we print the result using a "for var := range channel"
	*/
	for total := range sumNumbers(fibNumbers(numCounter)) {
		fmt.Print(total, "\n")
	}
}

// fibNumbers returns a send-only channel
func fibNumbers(qtt int) <-chan int {
	out := make(chan int) // output channel

	// This thread generates Fibonacci numbers
	go func() {
		// Initialization:
		newVal := 0
		lastVal := 1
		pastVal := 1

		// First two values
		out <- lastVal
		out <- pastVal

		// Generates new number by the sum of the last two values
		for i := 0; i < qtt-2; i++ { // qtt-2 : because we already have the first two values
			newVal = lastVal + pastVal
			pastVal = lastVal
			lastVal = newVal
			out <- newVal // goes to the channel
		}

		close(out) // closing output channel
	}()

	return out // returns the channel
}

// sumNumbers thakes a channel and sum every value stored in there, it also
// stores the result in a channel that is returned.
func sumNumbers(chNumbers <-chan int) <-chan int {
	out := make(chan int) // output channel

	// This thread just sum values stored in the "input" channel
	go func() {
		total := 0 // Where the result will be stored

		// Sum each value
		for aNum := range chNumbers {
			fmt.Print(aNum, " + ") // Prints the secuence
			total += aNum          // total = total + value
		}
		fmt.Print("\b = ")

		out <- total // the result
		close(out)   // closing output channel
	}()

	return out //returns the channel
}
