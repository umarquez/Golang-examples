package main

import (
	"encoding/json"
	"fmt"
)

// Country is a class designed to define Countrys,, even if those are islands.
type Country struct {
	Name       string
	Capital    string
	Population int
	Continent  string
	IsAnIsland bool
}

/**
  * ################################################
  * # Use the Unmarshal() method of json to work   #
	* # with strings or byteslices, to work with     #
  * # stereams use Decode() method instead.        #
  * ################################################
*/
func main() {
	var mx Country

	fmt.Println(mx.Name)
	fmt.Println(mx.Capital)
	fmt.Println(mx.Population)
	fmt.Println(mx.Continent)
	fmt.Println(mx.IsAnIsland)

	fmt.Print("\n\n===== [ space ] =====\n\n")

	values := `{
    "Name":"México",
    "Capital":"CDMX (DF)",
    "Population":122300000,
    "Continent":"America",
    "IsAnIsland": false,
    "president":"Enrique Peña Nieto"
  }`

	fmt.Println(values)

	bs := []byte(values)
	json.Unmarshal(bs, &mx) // This needs a pointer to 'mx' var to store values

	fmt.Print("\n===== [ space ] =====\n\n")

	fmt.Println(mx.Name)
	fmt.Println(mx.Capital)
	fmt.Println(mx.Population)
	fmt.Println(mx.Continent)
	fmt.Println(mx.IsAnIsland)
	fmt.Printf("%T\n", mx)
}
