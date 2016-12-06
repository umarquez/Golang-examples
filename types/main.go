package main

import "fmt"

/**
 * ############################################
 * # Custom type: person                      #
 * #                                          #
 * # Stores:                                  #
 * # - First Name as string                   #
 * # - Last name as string                    #
 * # - Age as int                             #
 * #                                          #
 * ############################################ */
type person struct {
	fName string
	lName string
	age   int
}

/**
 * ############################################
 * # Custom type: employee                    #
 * #                                          #
 * # Stores:                                  #
 * # - A var with person type                 #
 * # - Position as string                     #
 * # - Incomes as int                         #
 * #                                          #
 * ############################################ */
type employee struct {
	person
	position string
	incomes  int
}

// Function to get a Fullname of a person
func (p person) fullName() string {
	return p.fName + " " + p.lName
}

// Function to get a Fullname of an employee
func (p employee) fullName() string {
	return p.position + " " + p.fName + " " + p.lName
}

// Receive a pointer of an employee as argument
func gettingInfo(somebody *employee) {
	fmt.Println("")
	fmt.Println("===============")
	fmt.Println(somebody)                   // prints the content of the pointer location
	fmt.Printf("%T\n", somebody)            // Prints the type of "somebody"
	fmt.Println(&somebody)                  // Prints the pointer value (memmory address)
	fmt.Println(somebody.person.fullName()) // calling as "person" using pointer
}

func main() {
	me := employee{
		person: person{
			fName: "John",
			lName: "Doe",
			age:   30,
		},
		position: "Software developer",
		incomes:  1000,
	}

	myBoss := employee{
		person: person{
			fName: "Fulano",
			lName: "de Tal",
			age:   45,
		},
		position: "CTO",
		incomes:  5000,
	}

	fmt.Println(myBoss.fullName())    // Getting full name as employee (outer class overide inner class method)
	fmt.Println(me.person.fullName()) // Getting full name as person (specifying the class)

	gettingInfo(&me) // Passing as pointer of an employee
}
