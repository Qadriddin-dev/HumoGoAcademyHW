package main

import "fmt"

func main() {
	var firstNumber, secondNumber, thirdNumber int
	fmt.Print("Enter first number: ")
	fmt.Scan(&firstNumber)
	fmt.Print("Enter second number: ")
	fmt.Scan(&secondNumber)
	fmt.Print("Enter third number: ")
	fmt.Scan(&thirdNumber)

	isBetween := firstNumber < secondNumber && secondNumber < thirdNumber

	fmt.Println("Number B is between:", isBetween)
}
