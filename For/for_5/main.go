package main

import "fmt"

func main() {
	var priceKg float64

	fmt.Scan(&priceKg)

	for kg := 1; kg <= 10; kg++ {
		weight := float64(kg) / 10.0
		totalCost := priceKg * weight
		fmt.Printf("%.1f кг: %.2f\n", weight, totalCost)
	}
}
