package main

import "fmt"

func main() {
	var aNumber, bNumber float32
	fmt.Print("Напишите первое число: ")
	fmt.Scan(&aNumber)
	fmt.Print("Напишите второе число: ")
	fmt.Scan(&bNumber)

	average := (aNumber + bNumber) / 2

	fmt.Println("Среднее арифметическое:", average)
}
