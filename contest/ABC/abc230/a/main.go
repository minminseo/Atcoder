package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	if N >= 42 {
		fmt.Printf("AGC%.3d", N+1)
	} else {
		fmt.Printf("AGC%.3d", N)
	}
}
