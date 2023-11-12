package main

import (
	"fmt"
	"math"
)

func main() {
	var radius float64
	fmt.Print("Напишите радиус окружности: ")
	fmt.Scan(&radius)
	const pi = 3.14
	length := 2 * pi * radius
	square := pi * math.Pow(radius, 2)
	fmt.Println("Длина окружности:", length)
	fmt.Println("Площадь окружности:", square)
}
