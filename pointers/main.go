package main

import "fmt"

func main() {
	a := 43 // declares a variable

	fmt.Println(a)  // prints the var value
	fmt.Println(&a) // prints address (use "&" before var to get the address)

	var b *int = &a // declares b as an int pointer, then asign the a adress
	// the type of "b" is "*int" (int pointer)

	fmt.Println(b)  // prints de value of b (the memmory location of the "a" var)
	fmt.Println(*b) // get the value stored in the address referenced by "b"

	*b = 42        // store the value 42 into the memmory adrress refered by "b"
	fmt.Println(a) // a was changed by the last command
}
