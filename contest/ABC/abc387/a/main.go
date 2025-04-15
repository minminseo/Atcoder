package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(q1())
}

func q1() int {
	sc := bufio.NewScanner(os.Stdin)
	var m, n int

	sc.Scan()
	input := strings.Split(sc.Text(), " ")
	m, _ = strconv.Atoi(input[0])
	n, _ = strconv.Atoi(input[1])

	return (m + n) * (m + n)
}
