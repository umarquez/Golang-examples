package main

import (
	"fmt"
	"strconv"
)

func main() {
	var anInt = 80
	var aFloat = 3.141592
	var aString = "Hello world!"
	var aRune rune = 0x65
	var aSliceOfBytes = []byte{0x64, 0x64, 0x64}

	fmt.Printf("%T %v --> %T %v \n", anInt, anInt, float64(anInt), float64(anInt))
	fmt.Printf("%T %v --> %T %v \n", aFloat, aFloat, int(aFloat), int(aFloat))
	fmt.Printf("%T %v --> %T %v \n", anInt, anInt, string(anInt), string(anInt))
	fmt.Printf("%T %v --> %T %v \n", aString, aString, []rune(aString), []rune(aString))
	fmt.Printf("%T %v --> %T %v \n", aRune, aRune, string(aRune), string(aRune))
	fmt.Printf("%T %v --> %T %v \n", aSliceOfBytes, aSliceOfBytes, string(aSliceOfBytes), string(aSliceOfBytes))

	// ===== String Convert package:

	fmt.Print("\n==============================\n\n")

	// Functions from strconv library returns twoo values: the operation result
	// and an err object with <nil> value if everything ends ok
	fmt.Println(strconv.Itoa(anInt))
	fmt.Println(strconv.ParseBool("False"))
	fmt.Println(strconv.Atoi("357"))
	fmt.Println(strconv.ParseFloat("3.141592", 64))

}
