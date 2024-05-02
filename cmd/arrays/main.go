package main

import "fmt"

func main() {

	/*
		Arrays
		- Fixed size
		- Same type
		- Index access
		- Contiguous memory
	*/

	// Array declaration
	var intArr [3]string
	intArr[0] = "Meow"
	intArr[1] = "Meow"
	intArr[2] = "Nigga"

	fmt.Println(intArr[0])
	fmt.Println(intArr)
	fmt.Println(&intArr[0])

	// Array literal
	intArr2 := [3]int32{67, 68, 69}
	fmt.Println(intArr2)

	/*
		Slices
		- Dynamic size
		- Same type
		- Index access
		- Contiguous memory
	*/

	// Slice declaration
	var intSlice []int32
	intSlice = append(intSlice, 666)
	fmt.Println(intSlice[0])

	// Slice literal
	intSlice2 := []int32{1, 2, 3}
	fmt.Println(intSlice2)

	/*
		Maps
		- Key-Value pairs
	*/

	var newMap = map[string]uint{"Afzal": 22, "Ali": 20}
	fmt.Println(newMap["Afzal"])

	var age, ok = newMap["Nigga"]
	if ok {
		fmt.Printf("The age is %v", age)

	} else {
		fmt.Println("Invalid Name!")
	}

	// Loops

	for i, v := range intArr {
		fmt.Printf("Index: %v, Value: %v \n", i, v)
	}

	for name, age := range newMap {
		fmt.Printf("Name: %v, Age: %v \n", name, age)
	}

}
