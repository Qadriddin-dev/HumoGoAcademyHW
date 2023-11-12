package main

import "fmt"

func main() {
	var number int
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)

	if number > 0 {
		number += 1
	} else {
		number += 2
	}

	fmt.Println("Your number:", number)
}
