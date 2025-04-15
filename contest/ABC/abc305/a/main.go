package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a := 0
	if n%5 == 0 {
		fmt.Println(n)
	} else {
		a = n % 5
		if a <= 2 {
			fmt.Println(5 * (n / 5))
		} else {
			fmt.Println(5*(n/5) + 5)
		}
	}
}
