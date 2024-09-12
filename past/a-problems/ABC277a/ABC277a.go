package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	x, _ := strconv.Atoi(scanner.Text())

	p := make([]int, n)

	for i := 0; i < n; i++ {
		scanner.Scan()
		p[i], _ = strconv.Atoi(scanner.Text())
	}

	for i := 0; i < n; i++ {
		if p[i] == x {
			fmt.Println(i + 1)
		}
	}
}
