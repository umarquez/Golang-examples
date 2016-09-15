package outter

import "fmt"

/*
PrintVars function prints pkgVar who is declared at other.go file and a local
one who lives here in the function.
*/
func PrintVars() {
	test := "test content" // Local varialbe declaration
	fmt.Println(pkgVar)
	fmt.Println(test)
}
