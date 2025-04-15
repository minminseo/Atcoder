package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&A[i])
	}

	check := true
	for i := 1; i < N; i++ {
		if A[i] != A[0] {
			check = false
			break
		}
	}
	if check {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
