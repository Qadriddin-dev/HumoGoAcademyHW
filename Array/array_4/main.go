package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	count := 0

	for i := 1; i < n-1; i++ {
		if a[i] > a[i-1] && a[i] > a[i+1] {
			count++
		}
	}

	fmt.Println(count)
}
