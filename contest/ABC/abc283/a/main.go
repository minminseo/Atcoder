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

	fmt.Println(int(math.Pow((a), (b))))
}

/*
関数作るならこう
func pow(a, n int) int {
	result := 1
	for n > 0 {
		if n%2 == 1 {
			result = result * a
		}
		a = a * a
		n /= 2
	}
	return result
}
*/
