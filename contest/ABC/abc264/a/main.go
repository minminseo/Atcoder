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
	lrs := strings.Split(sc.Text(), " ")
	l, _ := strconv.Atoi(lrs[0])
	r, _ := strconv.Atoi(lrs[1])

	s := "atcoder"
	fmt.Println(s[l-1 : r])
}
