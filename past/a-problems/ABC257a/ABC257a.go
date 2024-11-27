package main

import (
	"fmt"
)

func main() {
	var n, x int
	m := 0
	fmt.Scan(&n, &x)
	alpha := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	if x%n == 0 {
		m = (x / n) - 1
	} else {
		m = x / n
	}
	fmt.Println(string(alpha[m]))

}
