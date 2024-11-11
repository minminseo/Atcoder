package main

import "fmt"

func main() {
	var N int
	var S string
	fmt.Scan(&N, &S)

	fmt.Println(S[N-1 : N])
}
