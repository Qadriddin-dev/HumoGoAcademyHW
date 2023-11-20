package main

import "fmt"

func main() {
	var aNumber, bNumber, sum int

	fmt.Scan(&aNumber)

	fmt.Scan(&bNumber)

	for i := aNumber; i <= bNumber; i++ {
		sum += i
	}
	fmt.Print(sum)
}
