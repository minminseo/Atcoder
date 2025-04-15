package main

import "fmt"

func main() {
	var A, B int
	fmt.Scan(&A, &B)

	if A%B == 0 {
		fmt.Println(A / B)
	} else {
		fmt.Println(A/B + 1)
	}
}
