package funcs_and_ctrls

import (
	"errors"
	"fmt"
)

func MessingWithFuncsAndCtrls() {
	fmt.Println("__________________________Exploring functions and control structurs__________________________")
	var printValue string = "some string to pass to print function"
	printMe(printValue)

	var numerator = 20
	var denominator = 5
	var result, remainder, err = integerDivision(numerator, denominator)
	// if err != nil {
	// 	fmt.Printf(err.Error())
	// } else if remainder == 0 {
	// 	fmt.Printf("The result of the integer division is: %v", result)
	// } else {
	// 	fmt.Printf("The remainder of the integer division is: %v", remainder)
	// }

	// NOTE: the above if else conditions can be written using a switch statement
	// NOTE: in Go, the break in switch statements are implied and thus redundant to inculude, but may be good practice to include them in some codebases for clarity, and readability
	switch {
	case err != nil:
		fmt.Printf(err.Error())
		break
	case remainder == 0:
		fmt.Printf("The result of the integer division is: %v", result)
		break
	default:
		fmt.Printf("The result of the integer division is: %v, and the remainder is: %v", result, remainder)
		break
	}

	// we can also create traditional/ conditional switches on a single variable
	switch remainder {
	case 0:
		fmt.Println("Division was exact")
		break
	case 1, 2:
		fmt.Println("Division was close")
		break
	default:
		fmt.Println("Division was not close")
	}
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

// Go allows us to return multiple values by specifying multiple return types(in parentheses) through our function signature
func integerDivision(numerator int, denominator int) (int, int, error) {
	// "error" is a type in Go, it is initialized to "nil"
	var err error
	if denominator == 0 {
		err = errors.New("Cannot Divide by Zero")
	}

	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}
