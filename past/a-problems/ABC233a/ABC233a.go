package main

import "fmt"

func main() {
	var X, Y int
	fmt.Scan(&X, &Y)

	var ans int
	for {
		if X >= Y {
			fmt.Println(ans)
			return
		} else {
			X += 10
			ans++
		}
	}
}
