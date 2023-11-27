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

	count := false

	for i := 0; i < n; i++ {
		if a[i] == x {
			count = true
		}
	}

	if count {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
