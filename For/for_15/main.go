package main

import (
	"fmt"
	"math"
)

func main() {
	var aNumber float64
	var nNumber float64

	fmt.Scan(&aNumber)
	fmt.Scan(&nNumber)

	result := math.Pow(aNumber, nNumber)

	fmt.Print(result)
}
