package main

import "fmt"

func main() {
	var aNumber, bNumber, square int

	fmt.Scan(&aNumber)

	fmt.Scan(&bNumber)

	for i := aNumber; i <= bNumber; i++ {
		square += i * i
	}
	fmt.Print(square)
}
