package main

import "fmt"

func main() {
	var N, L int
	fmt.Scanf("%d, %d", &N, &L)

	A := make([]int, N)
	count := 0

	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &A[i])

		if A[i] >= L {
			count++
		}
	}

	fmt.Println(count)

}
