package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	a := make([][]int, N)
	for i := range a {
		a[i] = make([]int, i+1)
		a[i][0], a[i][i] = 1, 1
		for j := 1; j < i; j++ {
			a[i][j] = a[i-1][j-1] + a[i-1][j]
		}
	}

	for _, r := range a {
		for _, num := range r {
			fmt.Print(num, " ")
		}
		fmt.Println()
	}
}
