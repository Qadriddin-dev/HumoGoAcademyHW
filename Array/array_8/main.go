package main

import "fmt"

func main() {
	var n, x int
	fmt.Scan(&n)

	a := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	fmt.Scan(&x)

	closest := a[0]

	minDiff := x - a[0]

	for i := 1; i < n; i++ {
		diff := x - a[i]
		if diff < minDiff {
			minDiff = diff
			closest = a[i]
		}
	}

	fmt.Println(closest)
}
