package main

import "fmt"

func main() {
	var priceKg float64

	fmt.Scan(&priceKg)

	for weight := 1.2; weight <= 2; weight += 0.2 {
		totalCost := priceKg * weight
		fmt.Printf("%.1f кг: %.2f\n", weight, totalCost)
	}
}
