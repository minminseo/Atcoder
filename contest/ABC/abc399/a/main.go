package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	s := scanner.Text()

	scanner.Scan()
	t := scanner.Text()

	count := 0
	for i := 0; i < n; i++ {
		if s[i] != t[i] {
			count++
		}
	}
	fmt.Println(count)
}
