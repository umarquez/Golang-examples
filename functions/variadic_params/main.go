package main

import "fmt"

func main() {
	r := average(12, 23, 34, 45, 56, 67, 78, 89, 90)
	fmt.Println(r)
}

// Function with a variadic argument, it is defined with the "..." before the arg type
func average(sf ...float64) float64 {
	fmt.Println(sf)        // display values of sf
	fmt.Printf("%T\n", sf) // display the type of sf
	var total float64

	for _, v := range sf { // For each element in sf (saved in v)
		total += v // sum each value
	}

	return total / (float64(len(sf))) // divided by the elements in sf, then return
}
