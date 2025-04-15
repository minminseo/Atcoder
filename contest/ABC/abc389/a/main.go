package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	fmt.Scan(&s)

	a, _ := strconv.Atoi(s[0:1])
	b, _ := strconv.Atoi(s[2:3])

	fmt.Println(a * b)
}
