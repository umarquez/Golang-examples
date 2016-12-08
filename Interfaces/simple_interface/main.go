package main

import "fmt"

// == Car object
type car struct {
	speedKmh int
}

func (c car) getDistance(minutes int) float32 {
	return float32(c.speedKmh) * (float32(minutes) / 60.0)
}

// == Bike object
type bike struct {
	speedMs int
}

func (b bike) getDistance(minutes int) float32 {
	return float32((b.speedMs * (minutes * 60)) / 1000)
}

/**
  * == Common interface:
	* Requieres that objects that implements has a function called
	* getDistance(x int) that returns a float32 as result.
	*
	* Both structs (bike and car) implements this interface because all they have
	* the required method with the according types.
*/
type vehicle interface {
	getDistance(minutes int) float32
}

// == A func that requieres a vehicle
func displayDistance(v vehicle, minutes int) {
	fmt.Println(v.getDistance(minutes), "km.")
}

/**
 * ################################
 */
const minutes int = 90 // = hour and a half

func main() {
	myBike := bike{5}     // spped: 5 m/s
	myMustang := car{200} // speed: 250 km/h

	displayDistance(myBike, minutes)
	displayDistance(myMustang, minutes)
}
