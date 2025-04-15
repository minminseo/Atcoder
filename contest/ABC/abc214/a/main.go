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

	if 1 <= n && n <= 125 {
		fmt.Println(4)
	} else if 126 <= n && n <= 211 {
		fmt.Println(6)
	} else {
		fmt.Println(8)
	}
}
