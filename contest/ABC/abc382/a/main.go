package main

import (
	"fmt"
)

func main() {
	var n, d int
	var s string
	fmt.Scanf("%d %d", &n, &d)
	fmt.Scanf("%s", &s)

	cnt := 0
	for _, c := range s {
		if string(c) == "." {
			cnt++
		}
	}

	if cnt == 0 {
		fmt.Print(d)
	} else {
		fmt.Print(d + cnt)
	}

}
