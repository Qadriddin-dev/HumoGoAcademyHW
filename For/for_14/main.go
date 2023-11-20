package main

import "fmt"

func main() {
	var nNumber int

	fmt.Scan(&nNumber)

	result := 0
	for i := 1; i <= 2*nNumber-1; i += 2 {
		result += i
	}

	fmt.Print(result)
}
