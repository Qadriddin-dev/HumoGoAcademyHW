package main

import "fmt"

func main() {
	var nNumber int

	fmt.Scan(&nNumber)

	if nNumber <= 0 {
		fmt.Println("Введите корректные значения для N")
		return
	}

	kNumber := 0
	result := 1

	for result*3 < nNumber {
		result *= 3
		kNumber++
	}

	fmt.Println(kNumber)
}
