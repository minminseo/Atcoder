package main

import (
	"fmt"
)

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)

	sum := a + b
	for i := 1; i <= n; i++ {
		var c int
		fmt.Scan(&c)
		if c == sum {
			fmt.Println(i)
			return
		}
	}
}
