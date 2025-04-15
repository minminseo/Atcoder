package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	xyns := strings.Split(sc.Text(), " ")
	x, _ := strconv.Atoi(xyns[0])
	y, _ := strconv.Atoi(xyns[1])
	n, _ := strconv.Atoi(xyns[2])
	if (x * 3) <= y {
		fmt.Println(x * n)
	} else {
		fmt.Println(y*(n/3) + x*(n%3))
	}
}
