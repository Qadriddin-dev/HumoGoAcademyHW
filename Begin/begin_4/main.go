package main

import "fmt"

func main() {
	var diameter float32
	fmt.Print("Напишите диаметр окружности: ")
	fmt.Scan(&diameter)
	const pi = 3.14
	length := pi * diameter
	fmt.Println("Длина окружности:", length)
}