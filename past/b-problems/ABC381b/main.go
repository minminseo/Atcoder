package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)

	if len(s)%2 != 0 {
		fmt.Println("No")
		return
	}

	cnt := make(map[rune]int)
	for _, c := range s {
		cnt[c]++
	}

	for _, v := range cnt {
		if v != 2 {
			fmt.Println("No")
			return
		}
	}

	for i := 0; i < len(s); i += 2 {
		if s[i] != s[i+1] {
			fmt.Println("No")
			return
		}
	}

	fmt.Println("Yes")
}
