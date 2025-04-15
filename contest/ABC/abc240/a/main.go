package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	ans := b - a
	if ans == 1 || ans == 9 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
