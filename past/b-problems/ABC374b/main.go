package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var sc = bufio.NewScanner(os.Stdin)
	sc.Scan()
	s := sc.Text()

	sc.Scan()
	t := sc.Text()

	if s == t {
		fmt.Println("0")
	} else {
		l := len(s)
		if len(s) > len(t) {
			l = len(t)
		}
		for i := 0; i < l; i++ {
			if s[i] != t[i] {
				fmt.Println(i + 1)
				return
			}
		}
		fmt.Println(l + 1)
	}
}
