package main

import "fmt"

func main() {
	var aNumber, bNumber, product int

	fmt.Scan(&aNumber)

	fmt.Scan(&bNumber)

	product = 1

	for i := aNumber; i <= bNumber; i++ {
		product *= i
	}
	fmt.Print(product)
}
