package main

import "fmt"

func main() {
	var nNumber int

	fmt.Scan(&nNumber)

	if nNumber <= 0 {
		fmt.Println("Введите корректные значения для N")
		return
	}

	isPower := false
	for currentN := nNumber; currentN > 0; currentN /= 3 {
		if currentN == 1 {
			isPower = true
			break
		} else if currentN%3 != 0 {
			break
		}
	}

	fmt.Println(isPower)
}
