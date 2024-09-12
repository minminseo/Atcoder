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

	h := make([]int, n)

	result := 0
	for i := 0; i < n; i++ {
		scanner.Scan()
		h[i], _ = strconv.Atoi(scanner.Text())
		if h[i] > result {
			result = h[i]
		}
	}

	for v := range h {
		if result == h[v] {
			fmt.Println(v + 1)
		}
	}

}
