package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	res := 0
	prev := 0
	for i := 0; i < n; i++ {
		var t, v int
		fmt.Scan(&t, &v)
		res -= t - prev
		if res < 0 {
			res = 0
		}
		res += v
		prev = t
	}

	fmt.Println(res)
}
