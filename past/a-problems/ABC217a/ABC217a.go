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
	sts := strings.Split(sc.Text(), " ")
	s, t := sts[0], sts[1]

	if s < t {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
