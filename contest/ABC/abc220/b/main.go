package main

import (
	"fmt"
	"strconv"
)

func main() {
	var k int
	var a, b string
	fmt.Scan(&k, &a, &b)
	a10, err := strconv.ParseInt(a, k, 64)
	if err != nil {
		panic(err)
	}
	b10, err := strconv.ParseInt(b, k, 64)
	if err != nil {
		panic(err)
	}
	fmt.Print(a10 * b10)
}
