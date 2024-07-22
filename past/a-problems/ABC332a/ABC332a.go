package main

import "fmt"

func main() {
	var N, S, K int
	fmt.Scanf("%d %d %d", &N, &S, &K)

	sum := 0
	for i := 1; i <= N; i++ {
		var P, Q int
		fmt.Scanf("%d %d", &P, &Q)
		sum += P * Q
	}

	if S <= sum {
		fmt.Println(sum)
	} else {
		fmt.Println(sum + K)
	}
}
