package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	s := sc.Text()

	if s[n-1] == 'o' {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
