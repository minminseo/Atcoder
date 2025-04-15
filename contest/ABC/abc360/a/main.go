package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)

	rs := "No"
	for i := 0; i < len(s); i++ {
		if s[i] == 'R' {
			rs = "Yes"
			break
		}
		if s[i] == 'M' {
			break
		}
	}
	fmt.Println(rs)
}
