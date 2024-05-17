package numbersStuff

import "fmt"

func MessingWithNumbers() {
	fmt.Println("__________________________Exploring Numbers__________________________")
	// int types are [int, int8, int16, int32, int64]
	var intNum int = 300
	// positive only ranges exist for all int types [uint, uint8, uint16, uint32, uint64]
	var intNum2 uint
	// float types are [float32, float64]
	var floatNum float32 = 83.625

	// in order to use different number types in an operation, one has to be cast to the type of the other.
	var result = floatNum + float32(intNum)
	fmt.Println("uint and float: ", intNum2+uint(result))

	var intNum3 int = 33
	var intNum4 int = 999
	// integer division results in an iteger that is rounded down
	// converting both ints to floats allows for un-rounded floating point precision.
	// NOTE use float64 to not round to nearest last decimal
	var intNumDivision = float32(intNum3) / float32(intNum4)

	fmt.Println("Integer division: ", intNumDivision)
}
