package main

import "fmt"

func main() {
	var nNumber int

	fmt.Scan(&nNumber)

	if nNumber <= 0 {
		fmt.Println("Введите корректные значения для N")
		return
	}

	result := 1.0

	for i := nNumber; i >= 1; i -= 2 {
		result *= float64(i)
	}

	fmt.Println(result)
}
