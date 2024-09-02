package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	a := make([]int, n)
	b := make([]int, m)

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	for i := 0; i < m; i++ {
		fmt.Scan(&b[i])
	}

	sum := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i+1 == b[j] {
				sum += a[i]
			}
		}
	}

	fmt.Println(sum)
}
