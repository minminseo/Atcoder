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
	k, _ := strconv.Atoi(scanner.Text())

	A := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		A[i], _ = strconv.Atoi(scanner.Text())
	}

	for i := n - k; i < n; i++ {
		fmt.Print(A[i], " ")
	}

	for i := 0; i < n-k; i++ {
		fmt.Print(A[i], " ")
	}
}
