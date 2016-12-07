package main

import (
	"encoding/json" // importing JSON library
	"fmt"
)

// Car is an exported custom type
type Car struct {
	Wheels      int // exported (starts with upper case)
	Doors       int
	ModelName   string
	IsHatchback bool `json:"Hatchback"` // using a tag to changes the ID on JSON mode
	ColorName   string
	KmsPerLiter float32 `json:"-"` // this tag will makes this field ignored on JSON mode
	alarmCode   string  // not exported (starts with lower case)
	ownerID     string
}

/**
  * ################################################
  * # Use the Marshal() method of json to work     #
	* # with strings or byteslices, to work with     #
  * # stereams use Encode() method instead.        #
  * ################################################
*/
func main() {
	// Car instace
	myCar := Car{
		Wheels:      5,
		Doors:       3,
		ModelName:   "Ford Focus Sport",
		IsHatchback: true,
		ColorName:   "Speedy red",
		KmsPerLiter: 21,
		alarmCode:   "0VKG5P7IPMHQRRBOIOD9", // Will not be exported in JSON mode
		ownerID:     "TNUUHES1TACHX74VRMID", // neither this
	}

	CarByteSlices, _ := json.Marshal(myCar) // converts to JSON in Byte Slices

	fmt.Println(string(CarByteSlices))
}
