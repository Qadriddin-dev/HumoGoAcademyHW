package main

import "fmt"

func main() {
	var aNumber, bNumber int

	fmt.Scan(&aNumber)

	fmt.Scan(&bNumber)

	count := 0

	for i := bNumber - 1; i > aNumber; i-- {
		fmt.Println(i)
		count++
	}
}
