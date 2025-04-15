package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	var S string
	fmt.Scan(&S)
	check := false

	for i := 0; i < N-2; i++ {
		if S[i] == 'A' && S[i+1] == 'B' && S[i+2] == 'C' {
			fmt.Println(i + 1)
			check = true
			return
		}
	}
	if !check {
		fmt.Println("-1")
	}
}
