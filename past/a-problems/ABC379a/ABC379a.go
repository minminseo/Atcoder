package main

import (
	"fmt"
	"strings"
)

func main() {
	var n string
	fmt.Scan(&n)
	strings := strings.Split(n, "")
	a := strings[0]
	b := strings[1]
	c := strings[2]

	fmt.Println(b+c+a, c+a+b)
}
