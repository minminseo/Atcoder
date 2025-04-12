package main

import "fmt"

func main() {
	var N, K int
	fmt.Scan(&N, &K)

	mod := int(1e9)
	A := make([]int, N+1)

	sum := 0
	for i := 0; i <= N; i++ {
		if i < K {
			A[i] = 1
		} else {
			A[i] = sum
		}
		sum = (sum + A[i]) % mod
		if i >= K {
			sum = (sum - A[i-K] + mod) % mod
		}
	}

	fmt.Println(A[N])
}
