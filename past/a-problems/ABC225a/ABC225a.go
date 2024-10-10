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

	if s[:1] == s[1:2] && s[:1] == s[2:] {
		fmt.Println(1)
	} else if s[:1] != s[1:2] && s[1:2] != s[2:] && s[:1] != s[2:] {
		fmt.Println(6)
	} else {
		fmt.Println(3)
	}
}
