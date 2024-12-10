package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	inputs := strings.Split(sc.Text(), " ")

	colorCount := make(map[string]int)
	for _, v := range inputs {
		colorCount[v]++
	}

	max := 0
	for _, count := range colorCount {
		max += count / 2
	}
	fmt.Println(max)
}
