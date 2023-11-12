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

	squareA := math.Pow(aNumber, 2)
	squareB := math.Pow(bNumber, 2)

	sum := squareA + squareB
	diff := squareA - squareB
	prod := squareA * squareB
	quotient := squareA / squareB

	fmt.Println("Сумма квадратов:", sum)
	fmt.Println("Разность квадратов:", diff)
	fmt.Println("Произведение квадратов:", prod)
	fmt.Println("Частное квадратов:", quotient)
}
