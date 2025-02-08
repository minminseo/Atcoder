package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	var h, hi, h2, h2i int
	for i := 0; i < n; i++ {
		var v int
		fmt.Scan(&v)
		if h < v {
			h2 = h
			h2i = hi
			h = v
			hi = i
			continue
		}
		if h2 < v {
			h2 = v
			h2i = i
		}
	}
	fmt.Print(h2i + 1)
}
