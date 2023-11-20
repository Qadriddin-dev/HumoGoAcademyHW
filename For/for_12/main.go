package main

import "fmt"

func main() {
	var nNumber int

	fmt.Scan(&nNumber)

	product := 1.0
	for i := 1; i <= nNumber; i++ {
		product += 1.1 * float64(nNumber)
	}

	fmt.Print(product)
}
