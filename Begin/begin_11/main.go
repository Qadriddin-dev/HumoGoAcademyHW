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

	modA := math.Abs(aNumber)
	modB := math.Abs(bNumber)

	sum := modA + modB
	diff := modA - modB
	prod := modA * modB
	quotient := modA / modB

	fmt.Println("Сумма модуля:", sum)
	fmt.Println("Разность модуля:", diff)
	fmt.Println("Произведение модуля:", prod)
	fmt.Println("Частное модуля:", quotient)
}
