package main

import "fmt"

func main() {
	var distanceInCentimeters float32

	fmt.Print("Введи расстояние в сантиметрах: ")
	fmt.Scan(&distanceInCentimeters)

	fullMeters := distanceInCentimeters / 100

	fmt.Println("Количество полных метров:", fullMeters)
}
