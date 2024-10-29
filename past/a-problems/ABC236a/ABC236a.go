package main

import (
	"fmt"
)

func main() {
	var s string
	var a, b int
	fmt.Scan(&s)
	fmt.Scan(&a, &b)
	s = s[:a-1] + string(s[b-1]) + s[a:b-1] + string(s[a-1]) + s[b:]
	fmt.Println(s)
}
