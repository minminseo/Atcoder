package main

import (
	"bufio"
	"fmt"
	"os"
)

var in = bufio.NewReader(os.Stdin)

func main() {
	var n, m int
	fmt.Fscan(in, &n, &m)

	s := makeSlice(n)
	t := makeSlice(m)

	match := false
	for i := 0; i < (n - m + 1); i++ {
		for j := 0; j < (n - m + 1); j++ {
			if s[i][j:j+m] == t[0] {
				match = true
				for k := 1; k < m; k++ {
					if s[i+k][j:j+m] != t[k] {
						match = false
					}
				}
				if match {
					fmt.Println(i+1, j+1)
					return
				}
			}
		}
	}

}

func makeSlice(n int) []string {
	a := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	return a
}
