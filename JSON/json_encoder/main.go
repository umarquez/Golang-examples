package main

import (
	"encoding/json"
	"os"
)

// Car is custom type
type Car struct {
	Wheels      int
	Doors       int
	ModelName   string
	IsHatchback bool
	ColorName   string
	KmsPerLiter float32
}

/**
 * ################################################
 * # Use the Enconde() method of an encoder to    #
 * # work with streams, if you want to work with  #
 * # byte slices or strings, use marshal instead. #
 * ################################################
 */
func main() {
	myCar := Car{
		Wheels:      5,
		Doors:       3,
		ModelName:   "Ford Focus Sport",
		IsHatchback: true,
		ColorName:   "Speedy red",
		KmsPerLiter: 21,
	}

	// NewEncoder() needs a file object with Write() method, os.Stdout is a file whith this method.
	json.NewEncoder(os.Stdout).Encode(myCar)

}
