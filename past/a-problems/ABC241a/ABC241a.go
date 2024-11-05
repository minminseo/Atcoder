package main

import "fmt"

func main() {
	var a = make([]int, 10)
	for i := 0; i < 10; i++ {
		fmt.Scan(&a[i])
	}
	var ans = 0
	for i := 0; i < 3; i++ {
		ans = a[ans]
	}
	fmt.Println(ans)
}
