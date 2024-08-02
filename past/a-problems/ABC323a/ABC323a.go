package main

import "fmt"

func main() {
	var S string
	fmt.Scan(&S)

	for i := 1; i < len(S); i += 2 {
		if S[i]%2 != 0 {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
