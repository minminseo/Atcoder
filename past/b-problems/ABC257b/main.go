package main

import (
	"fmt"
)

func main() {
	var n, k, q int
	fmt.Scan(&n, &k, &q)

	a := make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Scan(&a[i])
	}

	for i := 0; i < q; i++ {
		var l int
		fmt.Scan(&l)
		if a[l-1] < n {
			canMove := true
			for j := 0; j < k; j++ {
				if a[j] == a[l-1]+1 {
					canMove = false
					break
				}
			}
			if canMove {
				a[l-1]++
			}
		}
	}

	for i, v := range a {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
}
