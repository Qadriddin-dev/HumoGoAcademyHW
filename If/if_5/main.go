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
	negativeCount := 0

	if firstNumber > 0 {
		positiveCount++
	} else if firstNumber < 0 {
		negativeCount++
	}
	if secondNumber > 0 {
		positiveCount++
	} else if secondNumber < 0 {
		negativeCount++
	}
	if thirdNumber > 0 {
		positiveCount++
	} else if thirdNumber < 0 {
		negativeCount++
	}

	fmt.Println("Positive count number:", positiveCount)
	fmt.Println("Negative count number:", negativeCount)
}
