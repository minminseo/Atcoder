package main

import (
	"fmt"
)

func main() {
	var str string
	fmt.Scan(&str)
	var ans string
	for i := 0; i < 6; i += len(str) {
		ans += str
	}
	fmt.Println(ans)
}
