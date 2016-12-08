/**
	* ====================================================
  * TODO:
	* Use https://godoc.org/sort to sort the following (forward & reversed):
  *
	* (1)
	* type people []string
	* studyGroup := people{"Zeno", "John", "Al", "Jenny"}
	*
	* (2)
	* s := []string{"Zeno", "John", "Al", "Jenny"}
	*
	* (3)
	* n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	* ====================================================
*/

package main

import (
	"fmt"
	"sort"
)

// == people type definition
type people []string

// - implementing Interface interface:

// Len is the number of elements in the collection.
func (p people) Len() int {
	return len(p)
}

// Less reports whether the element with index i should sort before the element with index j.
// ==========
/*
// The long way:
func (p people) Less(i, j int) bool {
	var strLen int

	a := p[i] // first string
	b := p[j] // second string

	// getting the size of the smaller string
	if len(a) >= len(b) {
		strLen = len(b)
	} else {
		strLen = len(a)
	}

	// Compares its chars to define wich one goes first
	for x := 0; x < strLen; x++ {
		runeA := int(unicode.ToUpper(rune(a[x])))
		runeB := int(unicode.ToUpper(rune(b[x])))

		//fmt.Println(runeA, "==", runeB, "?") // <-- uncomment for debugging
		if runeA < runeB {
			return true
		} else if runeA > runeB {
			return false
		}
	}

	return false
}
*/
// ==========
// The short way:
func (p people) Less(i, j int) bool {
	reverse := false

	if reverse {
		return p[i] > p[j]
	}

	return p[i] < p[j]
}

// Swap swaps the elements with indexes i and j.
func (p people) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// == END people type

func main() {
	/**
	 * ###########################################################################
	 * # NOTE:                                                                   #
	 * #                                                                         #
	 * # Here's how to sort elements of a slice using a custom type and          #
	 * # sort.Interface (people struct), and using a internal custom type with   #
	 * # the sort.Interfce's functions (sort.StringSlice & sort.IntSlice).       #
	 * ###########################################################################
	 */
	studyGroup := people{"Zeno", "John", "Al", "Albert", "Jenny"}         // Using a custom type
	s := sort.StringSlice{"Zeno", "John", "Al", "Zoe", "Albert", "Jenny"} // Using sort.StringSlice istead of []string
	n := sort.IntSlice{7, 4, 8, 2, 9, 19, 12, 32, 3}                      // Using sort.IntSlice instead of []Int

	fmt.Println(studyGroup)             // Initial state
	sort.Sort(studyGroup)               // Sorting...
	fmt.Println(studyGroup)             // Second state
	sort.Sort(sort.Reverse(studyGroup)) // Reverse order
	fmt.Println(studyGroup)             // Final state

	fmt.Print("\n", "===== [ Space ] =====", "\n\n")

	fmt.Println(s)             // Initial state
	sort.Sort(s)               // Sorting...
	fmt.Println(s)             // Second state
	sort.Sort(sort.Reverse(s)) // Reverse order
	fmt.Println(s)             // Final state

	fmt.Print("\n", "===== [ Space ] =====", "\n\n")

	fmt.Println(n)             // Initial state
	sort.Sort(n)               // Sorting...
	fmt.Println(n)             // Second state
	sort.Sort(sort.Reverse(n)) // Reverse order
	fmt.Println(n)             // Final state
}
