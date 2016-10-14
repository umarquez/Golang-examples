package main

import "fmt"

func main() {
	data := []string{"John", "Doe", "Mary", "Ann"}
	fullName := getNames(data...) // Passing variadic args

	fmt.Println(fullName)
}

func getNames(names ...string) (out string) { // Recive variadic args and return a string
	for _, name := range names {
		out = fmt.Sprint(out, name, " ")
	}

	return
}
