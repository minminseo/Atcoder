package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	m, _ := strconv.Atoi(sc.Text())

	as := make([]int, m)
	for j := 0; j < m; j++ {
		sc.Scan()
		as[j], _ = strconv.Atoi(sc.Text())
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			sc.Scan()
			x, _ := strconv.Atoi(sc.Text())
			as[j] -= x
		}
	}

	for _, a := range as {
		if a > 0 {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
