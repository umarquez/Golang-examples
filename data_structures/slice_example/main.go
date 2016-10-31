package main

import "fmt"

func main() {

	mySlice := make([]int, 0, 5) // Declaring a "slice" with no items in it and capacity of 5

	fmt.Println("---------------")
	fmt.Println(mySlice)      // Printing content
	fmt.Println(len(mySlice)) // Printing number of elements
	fmt.Println(cap(mySlice)) // Printing capacity
	fmt.Println("---------------")

	for i := 0; i < 128; i++ {
		mySlice = append(mySlice, i) // adding "i" to the slice

		fmt.Println("Len:", len(mySlice), "Cap:", cap(mySlice), "Val:", mySlice[i])
	}
}
