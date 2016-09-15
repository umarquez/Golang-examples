package main

import "fmt"

/*
wrapper returns a function "func() int" when is called
*/
func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	x := 0

	increment := func() int { // Anonymous function
		x++
		return x
	}

	fmt.Println(increment()) // prints 1
	fmt.Println(increment()) // prints 2
	fmt.Println(increment()) // prints 3

	increment2 := wrapper()
	fmt.Println(increment2()) // prints 1
	fmt.Println(increment2()) // prints 2
}
