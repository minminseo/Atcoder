package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)
	var n int = len(s)
	var ans string
	for i := 0; i < n; i++ {
		if s[i] != '.' {
			ans += string(s[i])
		}
	}
	fmt.Println(ans)
}
