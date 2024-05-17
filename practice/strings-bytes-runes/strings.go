package strings_bytes_runes

import (
	"fmt"
	"strings"
)

func StringsBytesRunes() {
	fmt.Println("------------------Strings Bytes and Runes------------------")
	var myString = "Résumé"
	// when printing index 1 we get 195 as the code for the acute é because 2 byte characters
	//   need to be padded with some extra values and concatinated to fit the UTF8 representation
	//   and get the correct 233, the range key performs this computation under the hood.
	var indexed = myString[1]
	// alternitavely we can cast our string to an array of runes to get the same effect as the range keyword
	// NOTE: runes are essentiallly an alias for int32, and that strings in Go are immutable
	var someNewString = []rune("Résumé")
	fmt.Printf("%+v, %T, some new String: %+v\n", indexed, indexed, someNewString)

	/* yields this weird output with a missing index 2(because of the above mentioned UTF8 conversion)
	0 82
	1 233
	3 115
	4 117
	5 109
	6 233

	When working with strings in go, we're really working with an array of bytes
	all ASCII characters are 1 byte, and all else UTF8 chars for example are 2+ bytes
	*/
	for i, v := range myString {
		fmt.Println(i, v)
	}

	var strSlice = []string{"s", "u", "b", "s", "c", "r", "i", "b", "e"}
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	var catStr = strBuilder.String()
	fmt.Printf("cat str: %v, split version: %v\n", catStr, strings.Split(catStr, ""))

}
