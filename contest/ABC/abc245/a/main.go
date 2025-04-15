package main

import "fmt"

func main() {
	var A, B, C, D int
	fmt.Scan(&A, &B, &C, &D)

	if A < C || (A == C && B <= D) {
		fmt.Println("Takahashi")
	} else {
		fmt.Println("Aoki")
	}
}
