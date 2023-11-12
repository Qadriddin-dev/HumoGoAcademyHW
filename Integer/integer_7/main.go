package main

import "fmt"

func main() {
	var twoDigitNumber int

	fmt.Print("Введите двузначное число: ")
	fmt.Scan(&twoDigitNumber)

	tens := twoDigitNumber / 10
	units := twoDigitNumber % 10

	sum := tens + units
	product := tens * units

	fmt.Println("Сумма цифр:", sum)
	fmt.Println("Произведение цифр:", product)
}