package main

import (
	"fmt"
	"math"
)

func main() {
	var aNumber, bNumber float64
	fmt.Print("Напишите первое число: ")
	fmt.Scan(&aNumber)
	fmt.Print("Напишите второе число: ")
	fmt.Scan(&bNumber)

	average := math.Sqrt(aNumber * bNumber)

	fmt.Println("Среднее геометрическое:", average)
}
