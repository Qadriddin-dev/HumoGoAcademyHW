package main

import "fmt"

func main() {
	var firstNumber, secondNumber int
	fmt.Print("Enter first number: ")
	fmt.Scan(&firstNumber)
	fmt.Print("Enter second number: ")
	fmt.Scan(&secondNumber)

	if firstNumber > secondNumber {
		fmt.Println("First number is bigger")
	} else {
		fmt.Println("Second number is bigger")
	}

}
