package main

import "fmt"

func main() {
	var nNumber int

	fmt.Scan(&nNumber)

	result := 0.0
	sign := 1.0
	for i := 1; i <= nNumber; i++ {
		term := 1.0 + float64(i-1)*0.1
		result += sign * term
		sign *= -1
	}

	fmt.Print(result)
}
