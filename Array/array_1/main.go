package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	foundNegativePair := false

	for i := 1; i < n; i++ {
		if a[i] < 0 && a[i-1] < 0 {
			foundNegativePair = true
			break
		} else if a[i] > 0 && a[i-1] > 0 {
			foundNegativePair = true
			break
		}
	}

	if foundNegativePair {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
