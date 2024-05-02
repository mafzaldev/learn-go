package main

import (
	"errors"
	"fmt"
)

func main() {
	printMe("Hello World")
	var result, remainder, err = intDivision(10, 2)
	if err != nil {
		fmt.Println(err.Error())
	} else if remainder == 0 {
		fmt.Println("No remainder")
	}
	fmt.Printf("Result: %v, Remainder: %v", result, remainder)
}

func printMe(val string) {
	fmt.Println(val)
}

func intDivision(a int, b int) (int, int, error) {
	var err error
	if b == 0 {
		err = errors.New("Division by zero")
		return 0, 0, err
	}
	return a / b, a % b, err
}
