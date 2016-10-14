package main

import "fmt"

func main() {
	func() { // function declaration
		fmt.Println("This functions is executed...")
	}() // this "()" will call the function right after declaration
}
