package main

import (
	"fmt"
)

func main() {
	var s, t string
	fmt.Scan(&s, &t)

	pos := 0
	for _, cs := range s {
		for i := pos; i < len(t); i++ {
			if cs == rune(t[i]) {
				fmt.Printf("%d ", i+1)
				pos = i + 1
				break
			}
		}
	}
}
