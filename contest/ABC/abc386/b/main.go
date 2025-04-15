package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	sc.Scan()
	s := sc.Text()

	var count int
	for i := 0; i < len(s); i++ {
		count++
		if s[i] == '0' && i+1 < len(s) && s[i+1] == '0' {
			i++
		}
	}
	fmt.Println(count)
}
