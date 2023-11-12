package main

import "fmt"

func main() {
	var aSide, bSide int
	fmt.Print("Напишите A сторону прямоугольника: ")
	fmt.Scan(&aSide)
	fmt.Print("Напишите B сторону прямоугольника: ")
	fmt.Scan(&bSide)
	square := aSide * bSide
	fmt.Println("Площадь прямоугольника:", square)
	perimeter := (aSide + bSide) * 2
	fmt.Println("Периметр прямоугольника:", perimeter)
}
