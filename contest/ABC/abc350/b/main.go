package main

import (
	"fmt"
)

func main() {
	var N, Q, T int
	fmt.Scan(&N, &Q)

	teeth := make([]bool, N+1)
	for i := 0; i < Q; i++ {
		fmt.Scan(&T)
		teeth[T] = !teeth[T]
	}

	count := 0
	for i := 1; i <= N; i++ {
		if !teeth[i] {
			count++
		}
	}

	fmt.Println(count)
}
