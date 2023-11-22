package main

import "fmt"

func main() {
	var nNumber int

	fmt.Scan(&nNumber)

	sum := 1.0

	for i := 1; i <= nNumber; i++ {
		sum += 1.0 / float64(i)
	}
	fmt.Print(sum)
}
