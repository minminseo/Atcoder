package main

import (
	"fmt"
)

func main() {

	var t int
	fmt.Scan(&t)
	ans := f(f(f(t)+t) + f(f(t)))
	fmt.Println(ans)
}

func f(x int) (r int) {
	r = x*x + 2*x + 3
	return r
}
