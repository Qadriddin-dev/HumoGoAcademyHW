package main

import (
	"fmt"
	"math"
)

func main() {
	var nNumber int
	var aNumber float64

	fmt.Scan(&nNumber)
	fmt.Scan(&aNumber)

	result := math.Pow(aNumber, float64(nNumber))

	fmt.Print(result)
}
