package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	n, _ := strconv.Atoi(sc.Text())

	l := make([]string, 0, n)
	sc.Scan()
	l = append(l, sc.Text())

	for i := 1; i < n; i++ {
		sc.Scan()
		name := sc.Text()
		for j := 0; j < i; j++ {
			if name == l[j] {
				fmt.Println("Yes")
				return
			}
		}
		l = append(l, name)
	}
	fmt.Println("No")
}
