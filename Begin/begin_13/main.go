package main

import (
	"fmt"
	"math"
)

func main() {
	var r1, r2, pi float64

	fmt.Print("Напишите 1-ый радиус: ")
	fmt.Scan(&r1)

	fmt.Print("Напишите 2-ый радиус: ")
	fmt.Scan(&r2)

	pi = 3.14
	if r1 < r2 {
		fmt.Println("Ошибка r1 должен быть больше r2")
	} else {
		s1 := pi * math.Pow(r1, 2)
		s2 := pi * math.Pow(r2, 2)
		s3 := s1 - s2
		fmt.Println("Площадь S1 круга:", s1)
		fmt.Println("Площадь S2 круга", s2)
		fmt.Println("Внутренний радиус круга", s3)
	}
}
