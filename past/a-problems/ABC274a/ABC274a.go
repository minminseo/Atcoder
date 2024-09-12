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
	a, _ := strconv.ParseFloat(scanner.Text(), 64)

	scanner.Scan()
	b, _ := strconv.ParseFloat(scanner.Text(), 64)

	fmt.Printf("%.3f", b/a)
}
