package main

import "fmt"

func main() {
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			fmt.Println(x, " - ", y)
		}
	}
}
