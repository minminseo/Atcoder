package main

import (
	"fmt"
)

func main() {
	var v, a, b, c int
	fmt.Scan(&v, &a, &b, &c)

	for v >= 0 {
		v -= a
		if v < 0 {
			fmt.Print("F")
			return
		}
		v -= b
		if v < 0 {
			fmt.Print("M")
			return
		}
		v -= c
		if v < 0 {
			fmt.Print("T")
			return
		}
	}
}
