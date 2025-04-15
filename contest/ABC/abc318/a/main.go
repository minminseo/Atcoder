package main

import (
	"fmt"
)

func main() {
	var N, M, P int
	fmt.Scan(&N, &M, &P)

	result := M
	count := 0

	for N >= result {
		result += P
		count++
	}
	fmt.Println(count)
}
