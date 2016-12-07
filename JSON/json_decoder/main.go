package main

import (
	"encoding/json"
	"fmt"
	"strings"
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

func main() {
	var myCar Car

	/**
	 * NOTE: Do not put a "," at the last-item/field/value, it will break all
	 * the object.
	 *
	 * ex: {"a":1, "b":2, "c":3,} <-- that evil comma!
	 */
	strJSONCar := `{
    "Wheels":      5,
		"Doors":       3,
		"ModelName":   "Ford Focus Sport",
		"IsHatchback": true,
		"ColorName":   "Speedy red",
		"KmsPerLiter": 21
	}`

	rdr := strings.NewReader(strJSONCar)

	// rdr.WriteTo(os.Stdout) // <-- Uncomment for debugging

	// NewDecoder() needs a reader to take the JSON content and Decode() it using a pointer of an instance of Car
	err := json.NewDecoder(rdr).Decode(&myCar)

	// ----- Error control -----
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("========== [ Space ] ==========")
	fmt.Println("Wheels:", myCar.Wheels)
	fmt.Println("Doors:", myCar.Doors)
	fmt.Println("ModelName:", myCar.ModelName)
	fmt.Println("IsHatchback:", myCar.IsHatchback)
	fmt.Println("ColorName:", myCar.ColorName)
	fmt.Println("KmsPerLiter:", myCar.KmsPerLiter)
}
