package main

import "fmt"

func printSlice(slice []int) {
	for i, v := range slice {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
}

func main() {
	var n, p, q, r, s int
	fmt.Scan(&n, &p, &q, &r, &s)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	printSlice(a[:p-1])
	fmt.Print(" ")
	printSlice(a[r-1 : s])
	fmt.Print(" ")
	printSlice(a[q : r-1])
	fmt.Print(" ")
	printSlice(a[p-1 : q])
	fmt.Print(" ")
	printSlice(a[s:])
	fmt.Println()
}
