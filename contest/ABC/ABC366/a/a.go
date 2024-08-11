package main

import "fmt"

func main() {
	var n, t, a int
	fmt.Scan(&n, &t, &a)
	if t > a {
		t, a = a, t
	}
	ans := "No"
	if a > t+n-t-a {
		ans = "Yes"
	}
	fmt.Println(ans)
}
