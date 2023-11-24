package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	minElement := a[0]

	for i := 1; i < n; i++ {
		if a[i] < minElement {
			minElement = a[i]
		}
	}

	fmt.Println(minElement)
}
