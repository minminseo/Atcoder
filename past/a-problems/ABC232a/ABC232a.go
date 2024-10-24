package main

import (
	"fmt"
	"strconv"
)

func main() {
	var S string
	fmt.Scan(&S)

	a, _ := strconv.Atoi(S[:1])
	b, _ := strconv.Atoi(S[2:])
	fmt.Println(a * b)
}
