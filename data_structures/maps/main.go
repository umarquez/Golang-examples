package main

import "fmt"

func doTheMap(aMap map[int]string) {
	if val, exists := aMap[2]; exists {
		fmt.Println("The key/value exists (", exists, "): ", val)
		fmt.Println("This will be deleted.")
		delete(aMap, 2)
	} else {
		fmt.Println("The key/value does not exists (", exists, ").")
	}
}

func main() {
	// Declaring a map
	// var myMap = map[string]string{} // or
	// myMap := map[string]string{
	//   "key": "value",
	// 	"key2": "value 2"
	// } // or
	var myMap = make(map[string]string)
	myMap["key"] = "value"
	myMap["key2"] = "value 2"

	fmt.Println(myMap)
	fmt.Println("len: ", len(myMap))
	fmt.Printf("--------------------\n\n")

	var theMap = map[int]string{
		1: "a",
		2: "b",
		3: "c",
		4: "d",
		5: "e",
	}

	fmt.Println(theMap)
	// Because Maps and Slices are Referece Types, don't need to be passed byRef.
	doTheMap(theMap)
	fmt.Println(theMap)
}
