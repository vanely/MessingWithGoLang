package arrays

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

func CreatingAndSlicingArrays() {
	fmt.Println("------------------Creating And Slicing Arrays------------------")

	var arrOne = [...]int32{34, 6216, 755, 234, 66, 89, 234}
	arrOneSlice := arrOne[4:7]
	fmt.Println("slice of 'arrOne': ", arrOneSlice)

	var arrTwo [8]string = [8]string{"some", "value", "at", "this", "index", "six"}
	fmt.Println("string arr: ", arrTwo)

	someStr := "some string implicitly typed"
	strings.Split(someStr, "")

	someError := errors.New("an error message")
	fmt.Println(someError.Error())
	// slice
	arrThree := []error{errors.New("some error"), errors.New("some other error"), errors.New("more errorrrrr")}
	arrThreeCopy := make([]error, len(arrThree))
	copy(arrThreeCopy, arrThree)

	fmt.Printf("Copy of 'arrthree': %v\n", arrThreeCopy)

	fmt.Println("------------------Slice Speed Test [Uninitialized vs Initialized capacity]------------------")
	var n int = 1000000
	var testSliceOne = []uint32{}
	var testSliceTwo = make([]uint32, 0, n)

	fmt.Printf("total time without preallocation: %v\n", timeLoop(testSliceOne, n))
	fmt.Printf("total time with preallocation: %v\n", timeLoop(testSliceTwo, n))
}

func timeLoop(slice []uint32, n int) time.Duration {
	var t0 = time.Now()
	for len(slice) < n {
		slice = append(slice, 1)
	}
	return time.Since(t0)
}
