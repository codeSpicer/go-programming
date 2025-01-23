package main

import (
	"errors"
	"fmt"
)

func main() {

	const text string = "text to be printed"
	printMe(text)

	var num1, num2, err = intDivision(9, 4)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("The result of int division is %v , %v ", num1, num2)
}

func printMe(text string) {
	fmt.Println(text)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("Cannot divide by zero.")
		return 0, 0, err

	}
	return numerator / denominator, numerator % denominator, err
}
