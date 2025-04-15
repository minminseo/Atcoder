package main

import "fmt"

func main() {
	var S, T string

	fmt.Scan(&S)
	fmt.Scan(&T)

	if len(S) <= len(T) && S == T[:len(S)] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
