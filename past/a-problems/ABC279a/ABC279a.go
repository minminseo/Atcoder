package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	s := scanner.Text()

	count := 0
	for v := range s {
		if s[v] == 'v' {
			count++
		} else {
			count += 2
		}
	}
	fmt.Println(count)
}
