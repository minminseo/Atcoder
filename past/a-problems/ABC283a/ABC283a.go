package main

import (
	"bufio"
	"fmt"
	"math"
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

	fmt.Println(math.Pow(a, b))
}
