package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, x2 float64

	fmt.Print("Напишите 1-ый координат: ")
	fmt.Scan(&x1)

	fmt.Print("Напишите 2-ый координат: ")
	fmt.Scan(&x2)

	difference := math.Abs(x1) - math.Abs(x2)

	fmt.Println("Разница между координатами: ", difference)
}
