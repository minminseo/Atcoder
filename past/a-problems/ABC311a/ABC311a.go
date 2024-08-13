package main

import "fmt"

func main() {
	var n int
	var s string
	var a, b, c int

	fmt.Scan(&n)
	fmt.Scan(&s)

	for i := 0; i < n; i++ {
		switch s[i] {
		case 'A':
			a += 1
		case 'B':
			b += 1
		case 'C':
			c += 1
		}

		if (a > 0) && (b > 0) && (c > 0) {
			fmt.Println(i + 1)
			break
		}
	}
}
