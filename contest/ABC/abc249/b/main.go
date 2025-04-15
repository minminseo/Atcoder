package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	var uppercases = make(map[rune]int)
	var lowercases = make(map[rune]int)
	for _, c := range s {
		if c >= 'A' && c <= 'Z' {
			uppercases[c]++
		} else {
			lowercases[c]++
		}
	}
	if len(uppercases) == 0 {
		fmt.Println("No")
		return
	}
	if len(lowercases) == 0 {
		fmt.Println("No")
		return
	}
	for _, v := range uppercases {
		if v > 1 {
			fmt.Println("No")
			return
		}
	}
	for _, v := range lowercases {
		if v > 1 {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
