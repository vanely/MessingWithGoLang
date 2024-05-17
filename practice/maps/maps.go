package maps

import "fmt"

func MessingWithMaps() {
	fmt.Println("\n------------------Creating And Working With Maps------------------")
	// this type definition denotes that the keys will be of type "string" and the values will be of type "uint8"
	// NOTE: in go, a map will always return something even if the key doesn't exist because of how undefined entities defualt to their 0 value
	var newMap map[string]uint8 = make(map[string]uint8)
	fmt.Println("empty map: ", newMap)
	// alternitavely we can initialize a map with values like this
	var someOtherMap = map[string]uint8{"key1": 3, "key2": 93, "key3": 22}

	// when getting a value from a map 2 properties are returned, and value of they key we want, and the "ok" property that tells us whether the key exists in the map or not
	var val, ok = someOtherMap["key2"]
	if ok {
		fmt.Printf("getting map values: %v, and does it exist: %v\n", val, ok)
	} else {
		fmt.Println("Invalid name!")
	}

	// go has a built in delete for removing entries from a map, it takes the map, and the field to remove as args
	// this deletes by reference much like the implementation in javascript, so no return value is given
	delete(someOtherMap, "key3")
	fmt.Println("updated map: ", someOtherMap)

	// maps don't preserve order, so be careful when iteration may need to be sorted
	for key, val := range someOtherMap {
		fmt.Printf("key: %v, val: %v\n", key, val)
	}

	// this iteration pattern can also be used with arrays and slices where the first var is the index, and the second the value at index
	someArr := [5]int32{6, 234, 77, 13, 78}
	for i, val := range someArr {
		fmt.Printf("%v is at index %v\n", val, i)
	}

	// there are also no while loops in go, only for loops, which can be written in the form of while loops.
	arrLen := len(someArr)
	i := 0
	for i < arrLen {
		fmt.Println("(condition outside of loop) value at index: ", someArr[i])
		i++
	}

	// the looping condition can also be put inside of the loop
	j := 0
	for {
		if j == arrLen {
			break
		}
		fmt.Println("(condition inside of loop) value at index: ", someArr[j])
		j++
	}

	// or we could just use the traditional for loop syntax
	for k := 0; k < arrLen; k++ {
		fmt.Println("(traditional loop) value at index: ", someArr[k])
	}
}
