package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var s string
	fmt.Scan(&s)

	for i := 0; i < n; i++ {
		fmt.Print(string(s[i]))
		fmt.Print(string(s[i]))
	}
}
