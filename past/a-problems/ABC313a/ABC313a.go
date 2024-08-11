package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	P := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&P[i])
	}

	var ans int
	ans = 0
	for i := 1; i < n; i++ {
		if ans < P[i]-P[0]+1 {
			ans = P[i] - P[0] + 1
		}
	}

	fmt.Println(ans)
}
