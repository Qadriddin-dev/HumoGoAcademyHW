package main

import "fmt"

func main() {
	var number int
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)

	isOdd := number%2 == 0

	fmt.Println("Number is even:", isOdd)
}
