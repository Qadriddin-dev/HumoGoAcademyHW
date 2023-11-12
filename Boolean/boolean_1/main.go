package main

import "fmt"

func main() {
	var number int
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)

	isPositive := number > 0

	fmt.Println("Number is positive:", isPositive)
}
