package main

import "fmt"

func main() {
	var twoDigitNumber int

	fmt.Print("Введите двузначное число: ")
	fmt.Scan(&twoDigitNumber)

	ten := twoDigitNumber / 10
	unit := twoDigitNumber % 10

	fmt.Println("Десятки:", ten)
	fmt.Println("Единицы:", unit)
}
