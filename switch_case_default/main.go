package main

import "fmt"

func main() {
	text := "Tim"

	// Simple switch/case usage
	switch text {
	case "Tim":
		fmt.Println("Wassup Tim") // Case does not need a "break" because fallthrough is disabled by default
	case "Jenny":
		fmt.Println("Wassup Jenny")
	case "Marcus":
		fmt.Println("Wassup Marcus")
		fallthrough // Explicit fallthrough needed to enable this behavior
	case "Medhi":
		fmt.Println("Wassup Medhi")
		fallthrough
	case "Julian":
		fmt.Println("Wassup Julian")
	case "Sushant":
		fmt.Println("Wassup Sushant")
	}

	// multi-options switch
	switch text {
	case "Tim", "Jenny":
		fmt.Println("Wassup Tim, or, err, Jenny")
	case "Marcus", "Medhi":
		fmt.Println("Wassup Marcus, or, err, Medhi")
	case "Julian", "Sushant":
		fmt.Println("Wassup Julian, or, err, Sushant")
	default: // Default case, if none is true...
		fmt.Println("Hello x-tranger...")
	}
}
