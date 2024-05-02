package main

import "fmt"

func main() {
	var intNum int = 32767
	intNum = intNum + 1
	fmt.Println(intNum)

	var floatNum float64 = 3.14159
	fmt.Println(floatNum)

	var floatNum32 float32 = 3.14159
	var intNum32 int32 = 2

	var result float32 = float32(intNum32) + floatNum32
	fmt.Println("Result: ", result)

	var intNum1 int = 10
	var intNum2 int = 20

	fmt.Println("Sum: ", intNum1+intNum2)
	fmt.Println("Difference: ", intNum1-intNum2)
	fmt.Println("Product: ", intNum1*intNum2)
	fmt.Println("Quotient: ", intNum2/intNum1)
	fmt.Println("Remainder: ", intNum2%intNum1)

	var myString string = "Hello" + " " + "World!"
	fmt.Println(myString)

	var myBool bool = true
	fmt.Println(myBool)

	var1, var2 := 1, 2
	fmt.Println(var1, var2)

	const myConst int = 10
	fmt.Println(myConst)
}
