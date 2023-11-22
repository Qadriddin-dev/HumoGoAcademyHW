package main

import "fmt"

func main() {
	var nNumber int

	fmt.Scan(&nNumber)

	if nNumber <= 0 {
		fmt.Println("Введите корректные значения для N")
		return
	}

	kNumber := 1

	for kNumber*kNumber >= nNumber {
		kNumber++
	}

	fmt.Println(kNumber)
}
