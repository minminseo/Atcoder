package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	var a = make([]int, 10)
	for _, v := range s {
		dist := int(v - '0')
		a[dist] = 1
	}
	for i, v := range a {
		if v == 0 {
			fmt.Println(i)
		}
	}
}
