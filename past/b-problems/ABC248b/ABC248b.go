package main

import "fmt"

func main() {
	var A, B, K int

	fmt.Scanf("%d %d %d", &A, &B, &K)

	ans := 0
	for ans = 0; A < B; A *= K {
		ans++
	}
	fmt.Println(ans)
}
