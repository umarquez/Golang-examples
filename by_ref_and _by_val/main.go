package main

import "fmt"

func byVal(z int) {
	z = 0

	fmt.Printf("%p\t", &z)
	fmt.Println(z)
}

func byRef(z *int) {
	*z = 0

	fmt.Printf("%p\t", z)
	fmt.Println(*z)
}

func main() {
	x := 100

	fmt.Printf("%p\t", &x) // Location of x var
	fmt.Println(x)

	fmt.Println("- Calling byVal():")
	byVal(x)

	fmt.Printf("%p\t", &x) // Location of x var
	fmt.Println(x)

	fmt.Println("- Calling byRef():")
	byRef(&x)

	fmt.Printf("%p\t", &x) // Location of x var
	fmt.Println(x)
}
