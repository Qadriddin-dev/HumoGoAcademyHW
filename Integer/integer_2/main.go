package main

import "fmt"

func main() {
	var massInKilograms float32

	fmt.Print("Введите массу в килограммах: ")
	fmt.Scan(&massInKilograms)

	fullTons := massInKilograms / 1000

	fmt.Println("Количество полных тонн:", fullTons)
}
