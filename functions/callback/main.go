package main

import "fmt"

func printString(texts []string, callback func(string)) {
	for _, txt := range texts {
		callback(txt)
	}
}

func main() {
	simpleFunction := func(n string) { // Declaring a function
		fmt.Println(n)
	}

	// Call a function that calls a callback
	printString([]string{"a", "hola", "test", "John", "Doe", "etc..."}, simpleFunction)
}
