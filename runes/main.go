package main

import "fmt"

func main() {
	for i := 0x00; i <= 0xffffffff; i++ {
		fmt.Printf("%d - %#X - %v -%v \n", i, i, string(i), []byte(string(i)))
	}
}
