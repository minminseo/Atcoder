package main

import (
	"fmt"
)

func main() {
	var n int
	var s string

	fmt.Scan(&n)
	fmt.Scan(&s)

	good := false
	bad := false

	for i := 0; i < n; i++ {
		if s[i] == 'o' {
			good = true
		} else if s[i] == 'x' {
			bad = true
		}
	}

	if good && !bad {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
