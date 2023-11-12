package main

import "fmt"

func main() {
	var firstNumber, secondNumber int
	fmt.Print("Enter first number: ")
	fmt.Scan(&firstNumber)
	fmt.Print("Enter second number: ")
	fmt.Scan(&secondNumber)

	isInequality := firstNumber >= 0 || secondNumber < -2

	fmt.Println("Number is inequality:", isInequality)
}
