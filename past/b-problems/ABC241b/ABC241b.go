package main

import "fmt"

func main() {
	// Code for B - Pasta
	var N, M int
	fmt.Scan(&N, &M)

	var A, B int
	C := make(map[int]int, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&A)
		C[A]++
	}

	for i := 0; i < M; i++ {
		fmt.Scan(&B)
		C[B]--
		if C[B] < 0 {
			fmt.Println("No")
			return
		}
	}

	fmt.Println("Yes")
}
