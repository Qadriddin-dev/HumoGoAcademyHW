package main

import "fmt"

func main() {
	var nNumber, kNumber int

	fmt.Scan(&nNumber)

	fmt.Scan(&kNumber)
	if nNumber <= 0 || kNumber <= 0 {
		fmt.Println("Введите корректные значения для N и K.")
		return
	}

	quotient := 0
	for nNumber >= kNumber {
		nNumber -= kNumber
		quotient++
	}

	fmt.Println(quotient)
	fmt.Println(nNumber)
}
