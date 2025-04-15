package main

import (
	"fmt"
)

func main() {
	var s, t int
	fmt.Scan(&s, &t)

	cnt := 0
	for i := 0; i <= s; i++ {
		for j := 0; i+j <= s; j++ {
			for k := 0; i+j+k <= s; k++ {
				if i*j*k <= t {
					cnt++
				}
			}
		}
	}
	fmt.Println(cnt)
}
