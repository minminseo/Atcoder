package main

import (
	"fmt"
)

func main() {
	n, m, x, t, d := 0, 0, 0, 0, 0
	fmt.Scanf("%d %d %d %d %d", &n, &m, &x, &t, &d)

	if n >= m && x <= m {
		fmt.Print(t)
		return
	}

	fmt.Print(t - (x-m)*d)
}
