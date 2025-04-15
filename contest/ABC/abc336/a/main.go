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
	i := 0
	for n%2 == 0 {
		n /= 2
		i++
	}
	fmt.Println(i)
}
