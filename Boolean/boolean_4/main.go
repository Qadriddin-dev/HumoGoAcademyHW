package main

import "fmt"

func main() {
	var firstNumber, secondNumber int
	fmt.Print("Enter first number: ")
	fmt.Scan(&firstNumber)
	fmt.Print("Enter second number: ")
	fmt.Scan(&secondNumber)

	isInequality := firstNumber > 2 || secondNumber <= 3

	fmt.Println("Number is inequality:", isInequality)
}
