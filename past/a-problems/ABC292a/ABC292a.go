package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	upperStr := strings.ToUpper(s)
	fmt.Println(upperStr)
}
