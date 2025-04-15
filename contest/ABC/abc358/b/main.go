package main

import "fmt"

func main() {
	var n, a int
	fmt.Scanf("%d%d", &n, &a)
	var t = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&t[i])
	}
	var ans int
	for i := 0; i < n; i++ {
		if ans-t[i] < 0 {
			ans += t[i] - ans
		}
		ans += a
		fmt.Println(ans)
	}
}
