package main

import (
	"fmt"
	"strings"
)

func main() {
	var s = make([]string, 3)
	fmt.Scan(&s[0])
	fmt.Scan(&s[1])
	fmt.Scan(&s[2])

	var t string
	fmt.Scan(&t)

	var ans = make([]string, 0, 1000)
	for _, e := range t {
		v := int(e-'0') - 1
		ans = append(ans, s[v])
	}
	fmt.Println(strings.Join(ans, ""))
}
