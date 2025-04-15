package main

import (
	"fmt"
)

func main() {
	n := 0
	s := ""
	fmt.Scan(&n)
	fmt.Scan(&s)

	isSearch := false
	for _, v := range s {
		if v == '|' {
			isSearch = !isSearch
		}
		if isSearch && v == '*' {
			fmt.Print("in")
			return
		}
	}
	fmt.Print("out")
}
