package main

import "fmt"

func main() {
	var side int
	fmt.Print("Напишите сторону квадрата: ")
	fmt.Scan(&side)
	perimeter := side * 4
	fmt.Println("Периметр квадрата:", perimeter)
}
