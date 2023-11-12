package main

import "fmt"

func main() {
	var firstNumber, secondNumber int
	fmt.Print("Enter first number: ")
	fmt.Scan(&firstNumber)
	fmt.Print("Enter second number: ")
	fmt.Scan(&secondNumber)

	if firstNumber > secondNumber {
		firstNumber, secondNumber = secondNumber, firstNumber
	}

	fmt.Println("The new first number value:", firstNumber)
	fmt.Println("The new second number value:", secondNumber)
}
