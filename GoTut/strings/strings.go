package stringStuff

import (
	"fmt"
	sometest "go-tut/test"
	"unicode/utf8"
)

func MessingWithStrings() {
	fmt.Println("__________________________Exploring strings__________________________")
	// string can be created using the "string" type, and can be represented either with double quotes or back quotes. Single quotes are reserved for chars aka "rune"
	// when using double quotes, formatting a string requires string modifiers, such as \n, \t, etc...
	// when usingl back quotes the visual formatting is applied to the output string
	// EX:
	var myString string = "Some string with double quotes\n\tNew line with tab"
	var myOtherString string = `Some sting with back ticks
	New line with tab`

	fmt.Println(myString)
	fmt.Println(myOtherString)
	// import from test package
	// NOTE: using the built in len() function returns the amount of bytes in a string(all ascii characters/ alpha numerics are one byte so it can seem decieving.
	//       other special characters however are between 1 - 4 bytes so the length becomes inaccurate for character count
	var testString string = sometest.TestTheThing()
	fmt.Println("length of testString(bytes): ", len(testString))
	fmt.Println("length of testString(actual): ", utf8.RuneCountInString(testString))
}
