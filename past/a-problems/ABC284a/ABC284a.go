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

	s := make([]string, n)

	for i := 0; i < n; i++ {
		scanner.Scan()
		s[i] = scanner.Text()
	}

	for i := len(s) - 1; i >= 0; i-- {
		fmt.Println(s[i])
	}
}
