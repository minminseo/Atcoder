package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	s := make([]string, n)
	count := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&s[i])
		if s[i] == "For" {
			count++
		}
	}
	if 2*count > n {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
