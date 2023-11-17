package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64

	fmt.Print("Напишите A катет: ")
	fmt.Scan(&a)

	fmt.Print("Напишите B катет: ")
	fmt.Scan(&b)

	c := math.Sqrt(math.Pow(a, 2) + math.Pow(b, 3))
	P := a + b + c
	fmt.Println("Гипотенуза:", c)
	fmt.Println("Периметр:", P)
}
