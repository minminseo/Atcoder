package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n int
	_, _ = fmt.Scanln(&n)
	var sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	var currentHeight = 0
	for i := 0; i < n; i++ {
		sc.Scan()
		nextHeight, _ := strconv.Atoi(sc.Text())
		if nextHeight > currentHeight {
			currentHeight = nextHeight
		} else {
			break
		}
	}
	fmt.Println(currentHeight)
}
