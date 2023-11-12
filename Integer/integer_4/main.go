package main

import "fmt"

func main() {
	var aNumber, bNumber int

	fmt.Print("Введите значение A: ")
	fmt.Scan(&aNumber)

	fmt.Print("Введите значение B: ")
	fmt.Scan(&bNumber)

	countSegments := aNumber / bNumber

	fmt.Println("Количество отрезков B на отрезке A:", countSegments)
}