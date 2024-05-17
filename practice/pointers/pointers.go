package pointers

import "fmt"

// pointers are variables that store memory locations and not values
func MessingWithPointers() {
	fmt.Println("------------------Pointers------------------")
	// here we tell the compiler to give us a free memory location that's 32 bits wide(depends on OS)
	// NOTE: a pointer has to be assigned a memory location, or we'll get a nil pointer exception
	var p *int32 = new(int32)

	// to get the value at the pointers memory location we use a pointer dereference
	fmt.Printf("value at address of pointer: %v\n", *p)
	// we can also modify the value at the memory location in the same way we would reassign the value of a variable
	*p = 10
	fmt.Printf("modified memory location of pointer *p: %v\n", *p)

}
