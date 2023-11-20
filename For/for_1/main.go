package main

import "fmt"

func main() {
	var kNumber, nNumber int

	fmt.Scan(&kNumber)

	fmt.Scan(&nNumber)

	for i := 0; i < nNumber; i++ {
		fmt.Println(kNumber)
	}
}
