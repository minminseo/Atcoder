package main

import "fmt"

func main() {
	var N, K, A int
	fmt.Scan(&N, &K, &A)

	var ans int
	ans = (A + K - 1) % N
	if ans != 0 {
		fmt.Println(ans)
	} else {
		fmt.Println(N)
	}
}
