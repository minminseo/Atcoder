package main

import "fmt"

func main() {
	var n, d int
	fmt.Scan(&n, &d)
	t := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&t[i])
	}

	check := false
	for i := 0; i < n-1; i++ {
		if t[i+1]-t[i] <= d {
			fmt.Println(t[i+1])
			check = true
			break
		}
	}
	if !check {
		fmt.Println(-1)
	}
}
