package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < 998244353; i++ {
		if (n-i)%998244353 == 0 {
			fmt.Println(i)
		}
	}
}
