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
	a, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	b, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	c, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	d, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	e, _ := strconv.Atoi(scanner.Text())

	arr := make(map[int]bool)
	arr[a] = true
	arr[b] = true
	arr[c] = true
	arr[d] = true
	arr[e] = true

	fmt.Println(len(arr))
}
