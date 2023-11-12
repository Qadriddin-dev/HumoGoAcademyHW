package main

import (
	"fmt"
	"math"
)

func main() {
	var length float64
	fmt.Print("Напишите длину ребра куба: ")
	fmt.Scan(&length)
	volume := math.Pow(length, 3)
	square := 6 * math.Pow(length, 2)
	fmt.Println("Объем куба:", volume)
	fmt.Println("Площадь поверхности куба:", square)
}
