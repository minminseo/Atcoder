package main

import (
	"fmt"
)

func main() {

	var y int
	fmt.Scan(&y)

	x := y / 10
	if y > 0 || (y%10 == 0) {
		fmt.Println(x)
	} else {
		fmt.Println(x - 1)
	}
}
