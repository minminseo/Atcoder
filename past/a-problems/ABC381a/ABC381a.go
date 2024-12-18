package main

import (
	"fmt"
)

func main() {
	var n int
	var s string
	fmt.Scanf("%d\n%s", &n, &s)
	if n%2 == 0 || s[n/2] != '/' {
		fmt.Println("No")
		return
	}
	for i := 0; i < n/2; i++ {
		if s[i] != '1' || s[n-1-i] != '2' {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
	return
}
