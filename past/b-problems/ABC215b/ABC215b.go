package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	var v int = 1
	i := 0
	for {
		if v > n {
			fmt.Print(i - 1)
			return
		}
		v = v * 2
		i++
	}
}
