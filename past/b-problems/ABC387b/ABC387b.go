package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(q2())
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

func q2() int {
	sc := bufio.NewScanner(os.Stdin)

	sc.Scan()
	x, _ := strconv.Atoi(sc.Text())
	sum := 2025

	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			if i*j == x {
				sum -= i * j
			}
		}
	}
	return sum
}
