package main

import "fmt"

func main() {
	var firstNumber, secondNumber, numberIndex int
	fmt.Print("Enter first number: ")
	fmt.Scan(&firstNumber)
	fmt.Print("Enter second number: ")
	fmt.Scan(&secondNumber)

	if firstNumber < secondNumber {
		numberIndex = 1
	} else {
		numberIndex = 2
	}

	fmt.Println("Index of min number:", numberIndex)
}
