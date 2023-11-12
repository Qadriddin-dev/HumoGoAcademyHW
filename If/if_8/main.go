package main

import "fmt"

func main() {
	var firstNumber, secondNumber, larger, smaller int
	fmt.Print("Enter first number: ")
	fmt.Scan(&firstNumber)
	fmt.Print("Enter second number: ")
	fmt.Scan(&secondNumber)

	if firstNumber > secondNumber {
		larger = firstNumber
		smaller = secondNumber
	} else {
		larger = secondNumber
		smaller = firstNumber
	}

	fmt.Println("The largest number:", larger)
	fmt.Println("The smallest number:", smaller)
}
