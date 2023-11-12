package main

import "fmt"

func main() {
	var fileInBytes float32

	fmt.Print("Введите размер файла в байтах: ")
	fmt.Scan(&fileInBytes)

	fullKbytes := fileInBytes / 1024

	fmt.Println("Количество полных килобайтов:", fullKbytes)
}
