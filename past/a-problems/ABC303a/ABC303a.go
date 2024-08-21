package main

import (
	"fmt"
)

func main() {
	var n int
	var s, t string
	fmt.Scan(&n, &s, &t)
	for i := 0; i < n; i++ {
		if s[i] == t[i] {
			continue
		}
		if s[i] == 'l' {
			if t[i] == '1' {
				continue
			}
		} else if s[i] == '1' {
			if t[i] == 'l' {
				continue
			}
		} else if s[i] == 'o' {
			if t[i] == '0' {
				continue
			}
		} else if s[i] == '0' {
			if t[i] == 'o' {
				continue
			}
		}
		fmt.Println("No")
		return
	}
	fmt.Println("Yes")
}
