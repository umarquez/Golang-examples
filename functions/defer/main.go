package main

import "fmt"

func hello() {
	fmt.Print(" Hello")
}

func world() {
	fmt.Print(" World")
}

func noDefered() {
	world()
	hello()
}

func defered() {
	defer world() // This function call will be executed right before main ends.
	hello()

	// <-- imagine that here is a call like "world()", "defer" keyword do that.
}

func main() {
	noDefered()
	fmt.Println("")
	defered()
}
