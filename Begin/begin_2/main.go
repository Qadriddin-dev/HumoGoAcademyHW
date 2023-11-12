package main

import "fmt"

func main() {
	var side int
	fmt.Print("Напишите сторону квадрата: ")
	fmt.Scan(&side)
	square := side * side
	fmt.Println("Площадь квадрата:", square)
}
