package errorStuff

import (
	"errors"
	"fmt"
)

func MessingWithErrors() {
	fmt.Println("__________________________Exploring errors__________________________")
	err1 := errors.New("something went wrong and you suck for it")
	err2 := fmt.Errorf("some other error before: [%w]", err1)

	fmt.Println(err2)
	fmt.Println(errors.Unwrap(err2))
}
