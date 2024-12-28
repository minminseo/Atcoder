package main

import "fmt"

func main() {
	var S, T string
	fmt.Scan(&S, &T)

	if S == T {
		fmt.Println("Yes")
		return
	}

	var temp_S string

	for i := 0; i < len(S)-1; i++ {
		temp_S = S[:i] + S[i+1:i+2] + S[i:i+1] + S[i+2:]
		if temp_S == T {
			fmt.Println("Yes")
			return
		}
	}

	fmt.Println("No")
}
