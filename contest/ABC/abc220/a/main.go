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
	abcs := strings.Split(sc.Text(), " ")

	a, _ := strconv.Atoi(abcs[0])
	b, _ := strconv.Atoi(abcs[1])
	c, _ := strconv.Atoi(abcs[2])

	check := false
	for i := a; i <= b; i++ {
		if i%c == 0 {
			fmt.Println(i)
			check = true
			break
		}
	}

	if !check {
		fmt.Println(-1)
	}
}
