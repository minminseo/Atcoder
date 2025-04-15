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

	fmt.Println(Calculation(n))
}

func Calculation(n int) int {
	if n == 0 {
		return 1
	} else {
		result := n * Calculation(n-1)
		return result
	}
}
