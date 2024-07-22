package main

import "fmt"

func main() {
	var M, D int
	var y, m, d int

	fmt.Scanf("%d %d", &M, &D)
	fmt.Scanf("%d %d %d", &y, &m, &d)

	if D == d {
		if M == m {
			fmt.Println(y+1, 1, 1)
		} else {
			fmt.Println(y, m+1, 1)
		}
	} else {
		fmt.Println(y, m, d+1)
	}
}
