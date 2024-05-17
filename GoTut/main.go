// In go only one unique package name can exist per directory(different package names within the same directory are illegal)
package main // tells go compiler that this will be the entry point of the compiled go code.

import (
	"fmt"
	errorStuff "go-tut/errors"
	funcs_and_ctrls "go-tut/functions-and-control-structures"
	"go-tut/iterables"
	numberStuff "go-tut/numbers"
	stringStuff "go-tut/strings"
)

func main() {
	//_______________________________________________________
	// NUMBERS
	numberStuff.MessingWithNumbers()

	//_______________________________________________________
	// STRINGS
	fmt.Println("\n")
	stringStuff.MessingWithStrings()

	//_______________________________________________________
	// ERRORS
	fmt.Println("\n")
	errorStuff.MessingWithErrors()

	//_______________________________________________________
	// BOOLEANS & CONSTANTS
	// NOTE: every that is declared and not defined in Go defaults to its "0" value or its "falsy" value
	//       all numerics and runes default to 0, and strings default to empty string ""
	var someBool bool
	// constants are immutable, cannot be changed after being defined
	// constants also need to be initialized/ define, cannot be only declared
	const myConst string = "some constant value"

	fmt.Println("Default boolean value when not defined: ", someBool)
	fmt.Println("Value of a constant variable: ", myConst)

	//_______________________________________________________
	// FUNCTIONS & CONTROL STRUCTURES
	fmt.Println("\n")
	funcs_and_ctrls.MessingWithFuncsAndCtrls()

	//_______________________________________________________
	// FUNCTIONS & CONTROL STRUCTURES
	fmt.Println("\n")
	iterables.MessingWithIterables()
}
