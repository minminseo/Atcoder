package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	fmt.Scan(&N)

	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&A[i])
	}

	B := make([]int, N)
	copy(B, A)

	sort.Slice(B, func(i, j int) bool {
		return B[i] > B[j]
	})

	for i := 0; i < N; i++ {
		if A[i] == B[1] {
			fmt.Println(i + 1)
		}
	}
}
