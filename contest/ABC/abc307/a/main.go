package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	A := make([]int, 7*N)
	sum := 0

	for j := 0; j < 7*N; j++ {
		fmt.Scan(&A[j])
		sum += A[j]
		if (j+1)%7 == 0 {
			fmt.Print(sum, " ")
			sum = 0
		}
	}
}
