package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i <= n; i++ {
		for j := 0; j <= n; j++ {
			for k := 0; k <= n; k++ {
				if i+j+k <= n {
					fmt.Println(i, j, k)
				}
			}
		}
	}
}
