// mapの特性を利用

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var a, b, c, d, e int
	fmt.Fscan(in, &a, &b, &c, &d, &e)

	m := make(map[int]int)

	for _, v := range []int{a, b, c, d, e} {
		m[v]++
	}

	two := false
	three := false
	for _, v := range m {
		if v == 2 {
			two = true
		}
		if v == 3 {
			three = true
		}
	}

	if two && three {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
