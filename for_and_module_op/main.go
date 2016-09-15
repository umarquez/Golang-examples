package main

import "fmt"

func main() {
	for i := 1; i < 70; i++ {
		fmt.Printf("%d\t-\t", i)
		if i%2 == 1 {
			fmt.Println("Odd")
		} else {
			fmt.Println("Even")
		}
	}

	var num int

	fmt.Scan(&num)

	if num%2 == 1 {
		fmt.Println("Odd")
	} else {
		fmt.Println("Even")
	}
}
