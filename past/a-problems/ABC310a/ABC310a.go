package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, p, q int

	fmt.Scan(&n, &p, &q)

	d := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&d[i])
	}
	sort.Ints(d)

	if p < d[0]+q {
		fmt.Println(p)
	} else {
		fmt.Println(d[0] + q)
	}
}
