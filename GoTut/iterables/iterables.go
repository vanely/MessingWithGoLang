package iterables

import "fmt"

func MessingWithIterables() {
	fmt.Println("__________________________Exploring iterables__________________________")
	// square brackets followed by a type denotes that the varialbe will be an array of the corresponding types
	// 		when declaring the array, its max length can be specified by with a value inside of the brackets
	var intArr [3]int32
	newArr := [10]int32{343, 547, 2134, 98, 32, 99, 46, 122, 787, 8}
	// use the "[...]" in place of a numeric value to create an array of flexible size
	var implicitLengthArr = [...]int32{1, 78, 3, 4, 5, 655, 56, 3, 8, 5, 4, 62}
	intArr[1] = 123
	fmt.Println("intArr first val: ", intArr[0])
	// this notation of separated values by colon can be used to get sub arrays
	fmt.Println("intArr sub array: ", intArr[1:3])

	fmt.Println("newArr sub array: ", newArr[2:9])
	fmt.Println("implicitLengthArr random num: ", implicitLengthArr[9])

	// when we omit the array length value we create a slice
	// 		slices are powerful, and flexible arryas with some added functionality
	var intSlice = []int32{22, 33, 44, 55, 66, 77, 88}
	fmt.Println("slice without explicit length: ", intSlice)
}
