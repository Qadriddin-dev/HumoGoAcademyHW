package main

import "fmt"

func main() {
	var threeDigitNumber int

	fmt.Print("Введите трехзначное число: ")
	fmt.Scan(&threeDigitNumber)

	hundred := threeDigitNumber / 100

	fmt.Println("Первая цифра (сотни):", hundred)
}
