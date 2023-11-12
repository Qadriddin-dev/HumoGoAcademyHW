package main

import "fmt"

func main() {
	var threeDigitNumber int

	fmt.Print("Введите трехзначное число: ")
	fmt.Scan(&threeDigitNumber)

	units := threeDigitNumber % 10
	tens := (threeDigitNumber / 10) % 10

	fmt.Println("Последняя цифра:", units)
	fmt.Println("Средняя цифра:", tens)
}
