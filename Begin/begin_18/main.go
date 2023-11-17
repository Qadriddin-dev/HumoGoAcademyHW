package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, product float64

	fmt.Print("Напишите A координат: ")
	fmt.Scan(&a)

	fmt.Print("Напишите B координат: ")
	fmt.Scan(&b)

	fmt.Print("Напишите C координат: ")
	fmt.Scan(&c)

	AC := math.Abs(a - c)
	BC := math.Abs(b - c)

	product = AC * BC

	fmt.Println("AC координаты: ", AC)
	fmt.Println("BC координаты: ", BC)
	fmt.Println("Произведение  между координатами: ", product)
}
