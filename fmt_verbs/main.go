package main

import "fmt"

// Format verbs: https://golang.org/pkg/fmt/
func main() {
	var val = 42
	fmt.Printf("Dec \t %d \n", val)
	fmt.Printf("Bin \t %b \n", val)
	fmt.Printf("Chr \t %c \n", val)
	fmt.Printf("Oct \t %o \n", val)
	fmt.Printf("Hex \t %x \n", val)
	// use the "#" for adding the 0x prefix
	fmt.Printf("Hex \t %#x \n", val)
	fmt.Printf("Hex \t %#X \n", val)
}
