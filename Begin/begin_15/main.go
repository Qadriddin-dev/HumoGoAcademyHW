package main

import (
	"fmt"
)

func main() {
	var s, r, pi, l, d float64

	fmt.Print("Напишите площадь окружности: ")
	fmt.Scan(&s)

	pi = 3.14

	r = l / 2 * pi

	l = 2 * pi * r

	d = r * 2

	fmt.Println("Диаметр окружности: ", d)
	fmt.Println("Длина окружности: ", l)
}
