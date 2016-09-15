package main

/**
 * - Shorthand: it can be used only into a function.
 * - var: Could be use in any context and if the var is not initialized it will
 * have an empty value (0 for numbers, "" for strings, and so on)
 */

import "fmt"

// === Var method
// type declaration could be omitted, its value defines it.
var x string = "Hola mundo!" // <--
var y = "ASDF"
var i int
var o, p string = "value for \"o\"", "value for \"p\"" // multiple declaration

func main() {
	// === Shorthand method:
	a := 0
	b := "golang"
	c := 3.141592
	d := true

	fmt.Printf("%v ----- %T \n", a, a)
	fmt.Printf("%v ----- %T \n", b, b)
	fmt.Printf("%v ----- %T \n", c, c)
	fmt.Printf("%v ----- %T \n", d, d)
	fmt.Printf("%v ----- %T \n", x, x)
	fmt.Printf("%v ----- %T \n", y, y)
	fmt.Printf("%v ----- %T \n", i, i)
	fmt.Printf("%v ----- %T \n", o, o)
	fmt.Printf("%v ----- %T \n", p, p)
}
