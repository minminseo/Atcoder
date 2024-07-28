package main

import "fmt"

func main() {
	var N, X, S int
	result := 0
	fmt.Scanf("%d %d", &N, &X)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &S)
		if S <= X {
			result += S
		}
	}

	fmt.Println(result)
}
