package main

import "fmt"

func main() {
	var N, H, X int
	fmt.Scan(&N, &H, &X)
	P := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&P[i])
		if H+P[i] >= X {
			fmt.Println(i + 1)
			return
		}
	}
}
