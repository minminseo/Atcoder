package main

import (
	"fmt"
)

func main() {
	var N, p, A int
	fmt.Scan(&N, &p)

	for i := 1; i < N; i++ {
		fmt.Scan(&A)
		if p >= A {
			fmt.Println("No")
			return
		}
		p = A
	}

	fmt.Println("Yes")
}
