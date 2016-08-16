package main

import "fmt"

// Format verbs: https://golang.org/pkg/fmt/
func main() {
	for i := 64; i < 256; i++ {
		fmt.Printf("%d \t %b \t %#X \t %q\n", i, i, i, i)
	}
}
