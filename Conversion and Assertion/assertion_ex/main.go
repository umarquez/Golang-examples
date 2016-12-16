package main

import (
	"fmt"
)

func main() {

	var aString interface{} = "Lorem ipsum..." // This var is typed as an empty interface

	str, ok := aString.(string) // assertion to string type (returns the value of the assertion and true if everything is ok)

	fmt.Println(ok)
	fmt.Println(str)

	// -- Next code will not work because strExample is not an interface (uncomment to try)
	/**
		// Trying to use assertion on a string type var
	  var strExample = "Lorem ipsum..."
		str2, ok2 := strExample.(string) // --> invalid type assertion: strExample.(string) (non-interface type string on left)

		fmt.Println(ok2)
		fmt.Println(str2)
	*/
	// -- Wrong code ends here --
}
