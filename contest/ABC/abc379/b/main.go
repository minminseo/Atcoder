package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var to string
	var tn int
	var n int

	sc := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 4096)
	sc.Buffer(buf, math.MaxInt)

	sc.Scan()
	inputs := strings.Split(sc.Text(), " ")

	tn, _ = strconv.Atoi(inputs[0])
	n, _ = strconv.Atoi(inputs[1])

	sc.Scan()
	to = sc.Text()

	count := 0
	tCount := 0

	for i := 0; i < tn; i++ {
		if string(to[i]) == "O" {
			tCount++
		}
		if string(to[i]) == "X" {
			tCount = 0
		}

		if tCount == n {
			count++
			tCount = 0
		}
	}

	fmt.Print(count)
}
