package main

import "fmt"

func main() {
	var twoDigitNumber int

	fmt.Print("Введите двузначное число: ")
	fmt.Scan(&twoDigitNumber)

	tens := twoDigitNumber / 10
	units := twoDigitNumber % 10

	reversedNumber := units*10 + tens
	fmt.Println("Число после перестановки цифр:", reversedNumber)
}
