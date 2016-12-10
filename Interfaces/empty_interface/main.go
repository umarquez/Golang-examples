package main

import (
	"fmt"
)

// Vehicle definition
type Vehicle struct {
	name string
}

// Car definition:
type Car struct {
	Vehicle
}

func (c Car) String() string {
	return fmt.Sprintf("Hello! My name is "+c.name+" and I'm a [%T]"+`
          ____________
         / ___     ___ \
        / /   | | |   \ \
   ____/ /____| | |____\ ------\
  |   _____     |     _____    o)
 [---/     \---------/     \----]
     \_____/         \_____/
     `, c)
}

// END Car definition

// Bike definition:
type Bike struct {
	Vehicle
}

func (b Bike) String() string {
	return fmt.Sprintf("Hello! My name is "+b.name+" and I'm a [%T]"+`
                  ___
             ___     \
            ---|------\
        ___/_  |  / __|__
       /  /__\_o_/ /  |  \
       \_____/     \_____/
    `, b)
}

// END Bike definition

// Object interface definition:
type Object interface {
	String() string
}

// END Vehicle definition

//PrintObject is a function that requires an empty interface param,
//all types implements an empty interface so it could manage all objects
func PrintObject(o interface{}) {
	fmt.Println(o)
}

func main() {
	mustang := Car{
		Vehicle{
			name: "Mustang",
		},
	}

	bmx := Bike{
		Vehicle{
			name: "BMX",
		},
	}

	PrintObject(mustang)
	fmt.Println("=============================================")
	PrintObject(bmx)
}
