package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var m = make(map[string]int)
	for i := 0; i < n; i++ {
		var s string
		fmt.Scan(&s)
		m[s]++
	}
	var maxv int
	var ans string
	for e, v := range m {
		if maxv < v {
			maxv = v
			ans = e
		}
	}
	fmt.Println(ans)
}
