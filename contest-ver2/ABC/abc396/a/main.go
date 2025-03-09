package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&A[i])
	}

	for i := 0; i < N-2; i++ {
		if A[i] == A[i+1] && A[i] == A[i+2] {
			fmt.Println("Yes")
			return
		}
	}

	fmt.Println("No")
}
