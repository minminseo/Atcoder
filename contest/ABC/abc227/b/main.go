package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	s := make([]int, n)
	t := make([]bool, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&s[i])
	}

	for a := 1; a <= 1000; a++ {
		for b := 1; b <= 1000; b++ {
			for i, v := range s {
				if 4*a*b+3*a+3*b == v {
					t[i] = true
				}
			}
		}
	}

	var answer int
	for i := range t {
		if !t[i] {
			answer++
		}
	}
	fmt.Print(answer)
}
