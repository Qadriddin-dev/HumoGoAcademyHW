package main

import (
	"fmt"
	"math"
)

func main() {
	var l, r, pi, s float64

	fmt.Print("Напишите длинна окружности: ")
	fmt.Scan(&l)

	pi = 3.14

	r = l / 2 * pi

	s = pi * math.Pow(r, 2)
	fmt.Println("Радиус окружности: ", r)
	fmt.Println("Площадь окружности: ", s)
}
