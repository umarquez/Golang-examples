package main

import "fmt"

/*
Constant: a value that never changes...
In GO, do not use CAPITALIZATION in constants names except for ones that is
necesary, like Pi or MX...

UNTYPED constants:
A constant value that does not yet have a fixed type, that helps because the
type will be assigned when it will used.

ex.
Pi = 3.141592


TYPED constants:
A constant with a fixed type.

ex.
City string = "CDMX"
*/
const p string = "death & taxes"

// Multiple constants declaration:
const (
	_  = iota             // 0
	KB = 1 << (iota * 10) // 1 << (1 * 10) bit shifting 10 times
	MB = 1 << (iota * 10) // 1 << (2 * 10) bit shifting 20 times
	GB = 1 << (iota * 10) // 1 << (3 * 10) bit shifting 30 times
	TB = 1 << (iota * 10) // 1 << (4 * 10) bit shifting 40 times
	/*
	  iota: a small amount of...
	  Every time you invoque "iota" it increments in amount of 1
	*/

	// NOTE: https://godoc.org/math
	Pi = 3.141592
	MX = "México"
)

func main() {
	const q = 42

	fmt.Println("p - ", p)
	fmt.Println("q - ", q)
	fmt.Println()
	fmt.Println("=============================================================")
	fmt.Println("binary\t\t\t\t\t\tdecimal")
	fmt.Println("------\t\t\t\t\t\t-------")
	fmt.Printf("%b\t\t\t\t\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t\t\t\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t\t\t", GB)
	fmt.Printf("%d\n", GB)
	fmt.Printf("%b\t", TB)
	fmt.Printf("%d\n", TB)
}
