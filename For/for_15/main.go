package main

import (
	"fmt"
	"math"
)

func main() {
	var nNumber float64
	var aNumber float64

	fmt.Scan(&nNumber)
	fmt.Scan(&aNumber)

	result := math.Pow(aNumber, nNumber)

	fmt.Print(result)
}
