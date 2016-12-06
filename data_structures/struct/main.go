package main

import "fmt"

// Defining a new type wich has "struct" as undeline type
type aPerson struct {
	// adding fields to the new type
	fName   string "this is a tag..."
	lName   string `json:"last_name"` // tag of the JSON field's name
	age     int    `datastore:"-"`    // tag to avoid this field to be stored
	petName string
}

func main() {
	somebody := aPerson{"John", "Doe", 30, ""}          // new instance of aPerson type
	strager := aPerson{"Uriel", "Márquez", 32, "Tomas"} // new instance of aPerson type

	fmt.Print("Nombre: ", somebody.fName, "\nApellido: ", somebody.lName, "\nEdad: ", somebody.age, "\nMascota:", somebody.petName, "\n")
	fmt.Println("----------")

	somebody.petName = "Fido" // changging "petName" field's value

	fmt.Print("Nombre: ", somebody.fName, "\nApellido: ", somebody.lName, "\nEdad: ", somebody.age, "\nMascota:", somebody.petName, "\n")
	fmt.Println("----------")
	fmt.Print("Nombre: ", strager.fName, "\nApellido: ", strager.lName, "\nEdad: ", strager.age, "\nMascota:", strager.petName, "\n")
}
