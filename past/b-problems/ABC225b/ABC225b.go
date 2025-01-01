package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var N int
	fmt.Fscan(r, &N)

	x := make([]int, N)
	for i := 0; i < N-1; i++ {
		var a int
		var b int
		fmt.Fscan(r, &a, &b)
		x[a-1]++
		x[b-1]++
	}

	for _, v := range x {
		if v == (N - 1) {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}
