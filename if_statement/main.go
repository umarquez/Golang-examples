package main

import "fmt"

func main() {
	// Simple if example
	if true {
		fmt.Println("This always runs...")
	}

	var x = 5
	// if / else
	if x > 10 {
		fmt.Println("x is bigger than 10")
	} else {
		fmt.Println("x is not bigger enough")
	}

	var y = 10
	// if / else if / if
	if y < 10 {
		fmt.Println("y is smaller than 10")
	} else if y == 10 {
		fmt.Println("y is 10")
	} else {
		fmt.Println("y is bigger than 10")
	}

	// if with initialization
	if n := 666; n != 100 {
		fmt.Println("\"n\" (", n, ") is not 100")
	} else {
		fmt.Println("the \"n\" is 100")
	} // ----- The "n" scope ends right here

	// There is no more "n" available
	fmt.Println(n) // undefined var

}
