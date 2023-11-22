package main

import "fmt"

func main() {
	var aNumber, bNumber int

	fmt.Scan(&aNumber)

	fmt.Scan(&bNumber)
	if aNumber <= 0 || bNumber <= 0 || bNumber >= aNumber {
		fmt.Println("Введите корректные значения для A и B.")
		return
	}

	count := 0

	for aNumber >= bNumber {
		aNumber -= bNumber
		count++
	}

	fmt.Println(count)
}
