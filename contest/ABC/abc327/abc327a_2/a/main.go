package main

import "fmt"

func main() {
	var N int
	var S string
	fmt.Scan(&N, &S)

	yes := false
	for i := 0; i < N-1; i++ {
		if (S[i] == 'a' && S[i+1] == 'b') || (S[i] == 'b' && S[i+1] == 'a') {
			yes = true
			break
		}
	}

	if yes {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
