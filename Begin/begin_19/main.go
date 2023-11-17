package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, x2, y1, y2 float64

	fmt.Print("Напишите x1 сторону: ")
	fmt.Scan(&x1)

	fmt.Print("Напишите y1 сторону: ")
	fmt.Scan(&y1)

	fmt.Print("Напишите x2 сторону: ")
	fmt.Scan(&x2)

	fmt.Print("Напишите y2 сторону: ")
	fmt.Scan(&y2)

	a := math.Abs(x2 - x1)
	b := math.Abs(y2 - y1)

	perimeter := 2 * (a + b)

	area := a * b

	fmt.Println("Периметр прямоугольника: ", perimeter)
	fmt.Println("Площадь прямоугольника: ", area)
}
