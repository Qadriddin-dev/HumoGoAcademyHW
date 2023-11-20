package main

import "fmt"

func main() {
	var nNumber int

	fmt.Scan(&nNumber)

	sum := 0
	for i := 0; i <= nNumber; i++ {
		sum += (nNumber + i) * (nNumber + i)
	}

	sum /= 2
	fmt.Print(sum)
}
