package main

import "fmt"

func reverse(r []int) []int {
	n := len(r)
	reversed := make([]int, n)
	for i, v := range r {
		reversed[n-i-1] = v
	}
	return reversed
}

func main() {
	listA := make([]int, 100)
	result := make([]int, 100)
	for i := 0; i < 100; i++ {
		fmt.Scan(&listA[i])
		if listA[i] == 0 {
			result = listA[:i+1]
			break
		}
	}
	finalResult := reverse(result)

	for _, v := range finalResult {
		fmt.Println(v)
	}
}
