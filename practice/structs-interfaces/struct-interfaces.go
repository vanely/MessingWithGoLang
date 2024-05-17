package struct_interfaces

import (
	"fmt"
)

// structs are a way of defining your own types. EX:
type gasEngine struct {
	mpg       uint16
	gallons   uint16
	ownerInfo owner
}

type electricEngine struct {
	mpkwh uint16
	kwh   uint16
}

type owner struct {
	name string
}

// structs can have methods that have the context of the struct's instance
// these methods are like normal functions with the exception of a different type definition denoting that it belongs to the struct
func (e gasEngine) calculateMilesLeft() uint16 {
	return e.gallons * e.mpg
}

// and a function that's added to our 'electricEngine' struct for calculating miles left
func (e electricEngine) calculateMilesLeft() uint16 {
	return e.kwh * e.mpkwh
}

// *^* but what if we wanted a single milesLeft calculator function for both types
// we would then use an interface, defined in the same way:
type engine interface {
	calculateMilesLeft() uint16
}

// now we can make a function that takes an instance of a struct and call the method for calculating miles left that shouold be a part of these different structs
// we use the engine interface to denote that we'll be looking for a 'calculateMilesLeft' function within the instance of the struct that's passed in to the function
func canMakeIt(e engine, miles uint16) (string, uint16) {
	var milesLeft uint16 = e.calculateMilesLeft()
	var message string
	if miles <= milesLeft {
		message = "You can make it there!"
	} else {
		message = "Need to fuel up first!"
	}
	return message, milesLeft
}

func MessingWithStructsAndInterfaces() {
	fmt.Println("------------------Structs and Interfaces------------------")
	// now we can declare a variable that is the type of our struct(gasEngine)
	var myEngine gasEngine
	// we can define a variable with the struct literal syntax to initialize values for the struct type we're using
	var newEngine gasEngine = gasEngine{mpg: 25, gallons: 15, ownerInfo: owner{name: "Beatrix"}}
	var newElectricEngine electricEngine = electricEngine{mpkwh: 40, kwh: 10}
	// the field names can also be omitted and the order of the values are set to their respective fields
	var otherEngine gasEngine = gasEngine{32, 20, owner{"Joao"}}
	// we can also define one field at a time, and use other structs as field types within structs
	fmt.Printf("Initial values of 'myEngine': %+v, and 'newEngine': %+v\n", myEngine, newEngine)
	fmt.Printf("omitted field names: %+v\n", otherEngine)

	// we can create anonymous structs. Structs that are created with variable declaration, these are not reusable
	var specsAnonymStruct = struct {
		torque     uint16
		horsePower uint16
		curbWeight uint16
		driveTrain string
	}{210, 240, 3200, "FWD"}

	fmt.Printf("car specs: %+v\n", specsAnonymStruct)
	fmt.Printf("%+v\n", specsAnonymStruct)
	fmt.Printf("miles left in car: %v, belonging to: %v\n", newEngine.calculateMilesLeft(), newEngine.ownerInfo.name)
	// *^* using calculation function of any struct through more generic interface type
	gasMessage, gasMilesLeft := canMakeIt(newEngine, 400)
	electricMessage, electricMilesLeft := canMakeIt(newElectricEngine, 300)
	fmt.Printf("using interface on gasEngine (message: %v, milesLeft: %v)\n", gasMessage, gasMilesLeft)
	fmt.Printf("using interface on electricEngine (message: %v, milesLeft: %v)\n", electricMessage, electricMilesLeft)
}
