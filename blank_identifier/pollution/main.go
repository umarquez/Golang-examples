package main

import "fmt"

func main() {
	// You must use everything you declare, elsewhere you are polluting your code.
	var a = 1
	var b = 2 // Noy used...

	fmt.Println(a)
	_ = b // avoiding "declared and not used" error
}
