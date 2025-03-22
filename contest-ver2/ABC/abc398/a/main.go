package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())
	var s string

	if N%2 == 0 {
		s = strings.Repeat("-", N/2-1) + "==" + strings.Repeat("-", N/2-1)
	}
	if N%2 != 0 {
		s = strings.Repeat("-", N/2) + "=" + strings.Repeat("-", N/2)
	}

	fmt.Println(s)
}
