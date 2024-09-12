package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	s := scanner.Text()

	check := false
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == 'a' {
			fmt.Println(i + 1)
			check = true
			break
		}
	}

	if !check {
		fmt.Println(-1)
	}
}
