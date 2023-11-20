package main

import "fmt"

func main() {
	var priceKg float64

	fmt.Scan(&priceKg)

	for kg := 1; kg <= 10; kg++ {
		totalCost := priceKg * float64(kg)
		fmt.Printf("%d кг: %.2f\n", kg, totalCost)
	}
}
