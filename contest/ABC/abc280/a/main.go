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
	h, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	w, _ := strconv.Atoi(scanner.Text())

	s := make([]string, h)

	for i := 0; i < h; i++ {
		scanner.Scan()
		s[i] = scanner.Text()
	}

	count := 0

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if s[i][j] == '#' {
				count++
			}
		}
	}
	fmt.Println(count)
}
