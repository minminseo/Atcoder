package main

import (
	"fmt"
)

func main() {
	s := ""
	fmt.Scan(&s)

	count := make(map[string]int, len(s))
	for _, v := range s {
		count[string(v)]++
	}

	for k, v := range count {
		if v == 1 {
			fmt.Print(k)
			return
		}
	}
	fmt.Print(-1)
}
