package main

import "fmt"

func main() {
	var arrayData [7]string // Array has a defined number of elements inside "[]"
	var sliceData []string  // Slice hasn't a defined number of elements

	/*
	 * Array: a static list of items, it can't be resized
	 * Slice: a dynamic list of items, it can be resized
	 **/

	arrayData = [7]string{
		"hello",
		"world",
		"asdf",
		"1q2w3e",
		"qwerty",
		"abc",
		"def",
	}

	sliceData = []string{"a1", "b2", "c3", "d4", "e5", "f6", "g7", "h8", "i9", "j10"}

	fmt.Println(arrayData)
	fmt.Println(len(arrayData))
	fmt.Printf("%T\n\n", arrayData)

	fmt.Println(sliceData)
	fmt.Println(len(sliceData))
	fmt.Printf("%T\n", sliceData)
	fmt.Println(sliceData[2:])  // From index 2 to the end
	fmt.Println(sliceData[5:7]) // From index 2 to the index before index 7
	fmt.Println(sliceData[:2])  // From the beginning to the index before index 2
}
