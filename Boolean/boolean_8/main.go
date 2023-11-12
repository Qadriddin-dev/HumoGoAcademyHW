package main

import "fmt"

func main() {
	var firstNumber, secondNumber int
	fmt.Print("Enter first number: ")
	fmt.Scan(&firstNumber)
	fmt.Print("Enter second number: ")
	fmt.Scan(&secondNumber)

	bothOdd := firstNumber%2 != 0 && secondNumber%2 != 0

	fmt.Println("Both number is odd:", bothOdd)
}
