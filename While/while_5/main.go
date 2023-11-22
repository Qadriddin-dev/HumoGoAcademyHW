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
	for currentN := nNumber; currentN > 1; currentN /= 2 {
		kNumber++
	}

	fmt.Println(kNumber)
}
