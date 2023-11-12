package main

import "fmt"

func main() {
	var aLength, bLength, cLength int
	fmt.Print("Напишите длину A ребра параллелепипеда: ")
	fmt.Scan(&aLength)
	fmt.Print("Напишите длину B ребра параллелепипеда: ")
	fmt.Scan(&bLength)
	fmt.Print("Напишите длину C ребра параллелепипеда: ")
	fmt.Scan(&cLength)
	volume := aLength * bLength * cLength
	square := 2 * (aLength*bLength + bLength*cLength + aLength*cLength)
	fmt.Println("Объем куба:", volume)
	fmt.Println("Площадь поверхности параллелепипеда:", square)
}
