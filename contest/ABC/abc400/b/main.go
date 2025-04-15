package main

import (
	"fmt"
)

func main() {
	var N, M int
	fmt.Scan(&N, &M)

	const limit = 1000000000
	sum := 0
	power := 1

	for i := 0; i <= M; i++ {
		sum += power
		if sum > limit {
			fmt.Println("inf")
			return
		}
		power *= N
	}

	fmt.Println(sum)
}
