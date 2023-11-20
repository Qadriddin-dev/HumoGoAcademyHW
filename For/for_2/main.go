package main

import "fmt"

func main() {
	var aNumber, bNumber int

	fmt.Scan(&aNumber)

	fmt.Scan(&bNumber)

	count := 0

	for i := aNumber; i <= bNumber; i++ {
		fmt.Println(i)
		count++
	}
}
