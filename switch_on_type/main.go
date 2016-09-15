package main

import "fmt"

type contact struct {
	greeting string
	name     string
}

func switchOnType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case contact:
		fmt.Println("contact")
	default:
		fmt.Println("unknow")
	}
}

func main() {
	var t = contact{"Hola", "Mundo"}

	switchOnType(7)
	switchOnType("Hola...")
	switchOnType(t)
	switchOnType(3.1416)
}
