package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	var a = make([]int, n)
	var b = make([]int, k)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	for i := 0; i < k; i++ {
		fmt.Scan(&b[i])
		b[i]--
	}

	var max int
	for _, v := range a {
		if max < v {
			max = v
		}
	}
	for _, v := range b {
		if max == a[v] {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}
