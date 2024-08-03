// わからんかった

package main

import (
	"fmt"
)

func main() {
	var N, M int
	fmt.Scan(&N, &M)

	A := make([]int, N)
	sum := 0

	for i := 0; i < N; i++ {
		fmt.Scan(&A[i])
		sum += A[i]
	}

	if sum < M {
		fmt.Println("infinite")
	} else {

	}
}
