package main

import (
	"fmt"
)

func main() {
	var n, w int
	fmt.Scan(&n)
	rank := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		var s string
		fmt.Scan(&s)
		w = 0
		for _, c := range s {
			if c == 'o' {
				w++
			}
		}
		rank[w] = append(rank[w], i)
	}
	for i := n; i >= 0; i-- {
		for _, p := range rank[i] {
			fmt.Print(p, " ")
		}
	}
}
