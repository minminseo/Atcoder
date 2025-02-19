package main

import (
	"fmt"
)

func main() {
	n := 0
	fmt.Scan(&n)

	names := make([]string, n)
	ages := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&names[i])
		fmt.Scan(&ages[i])
	}

	min := ages[0]
	idx := 0
	for i, v := range ages {
		if min > v {
			min = v
			idx = i
		}
	}

	for i := 0; i < n; i++ {
		fmt.Println(names[idx])
		if idx == n-1 {
			idx = 0
			continue
		}
		idx += 1
	}
}
