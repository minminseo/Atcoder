// https://atcoder.jp/contests/abc385/tasks/abc385_a
package main

import (
	"fmt"
)

func main() {
	var (
		a int
		b int
		c int
	)

	fmt.Scan(&a, &b, &c)

	var result string
	if a+b == c {
		result = "Yes"
	} else if b+c == a {
		result = "Yes"
	} else if a+c == b {
		result = "Yes"
	} else if a == b && a == c {
		result = "Yes"
	} else {
		result = "No"
	}

	fmt.Println(result)
}
