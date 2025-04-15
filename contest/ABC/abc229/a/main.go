package main

import "fmt"

func main() {
	var s1, s2 string
	fmt.Scan(&s1, &s2)
	if s1 == "#." && s2 == ".#" {
		fmt.Print("No")
		return
	}
	if s1 == ".#" && s2 == "#." {
		fmt.Print("No")
		return
	}
	fmt.Print("Yes")
}
