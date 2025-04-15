package main

import "fmt"

func main() {
	var L1, R1, L2, R2, count int
	fmt.Scanf("%d %d %d %d", &L1, &R1, &L2, &R2)

	for i := 0; i < R2; i++ {
		if L1 <= i && i < R1 && L2 <= i && i <= R2 {
			count++
		}
	}

	fmt.Println(count)
}
