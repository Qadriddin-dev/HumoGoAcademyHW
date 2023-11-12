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

	positiveCount := 0

	if firstNumber > 0 {
		positiveCount++
	}
	if secondNumber > 0 {
		positiveCount++
	}
	if thirdNumber > 0 {
		positiveCount++
	}

	fmt.Println("Positive count number:", positiveCount)
}
