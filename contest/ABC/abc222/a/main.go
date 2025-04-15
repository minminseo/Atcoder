package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	s := sc.Text()

	if 4-len(s) > 0 {
		fmt.Println(strings.Repeat("0", 4-len(s)) + s)
	} else {
		fmt.Println(s)
	}

}
