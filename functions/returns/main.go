package main

import "fmt"

func main() {
	fmt.Println(helloWorld()) // call with no args.
	fmt.Println(greet("John ", "Doe"))

	x, y := names("John", "Doe") // multiple values returned
	fmt.Println(x)
	fmt.Println(y)
}

// Normal function with no arguments.
func helloWorld() string {
	return "Hello World..." // A "traditional" return.
}

// This function will return the named var "s"
func greet(fname string, lname string) (s string) {
	s = fmt.Sprint(fname, lname) // for strings cocatenation only.
	return
}

// This function returns multiple values when its called
func names(fname, lname string) (string, string) {
	return fmt.Sprint("1.- ", fname, " ", lname), fmt.Sprint("2.- ", lname, " ", fname) // Here...
}
