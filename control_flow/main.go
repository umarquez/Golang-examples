package main

import "fmt"

func main() {
	/* ===== FOR ===== */
	// Like a C/Java for:
	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}

	x := 0
	// "While" style
	for x <= 100 {
		fmt.Println(x)
		x++
	}

	y := 0
	// Infinite loop
	for {
		fmt.Println("Doing this for ever...")

		// Ok, maybe only 11 times...
		if y >= 10 {
			break
		}
		y++
	}

	for n := 0; n <= 100; n++ {

		if n%2 == 0 { // even or odd?
			continue
		} else if n >= 100 { // over the limit?
			break
		} else { // else
			fmt.Println(n)
		}
	}
}
