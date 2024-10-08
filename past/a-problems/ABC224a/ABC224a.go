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

	if s[len(s)-2:] == "er" {
		fmt.Println("er")
	} else {
		fmt.Println("ist")
	}
}
