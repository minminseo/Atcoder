package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()

	nums := make([]bool, 2001, 2001)

	for _, a := range strings.Split(s, " ") {
		ai, _ := strconv.Atoi(a)
		nums[ai] = true
	}

	var result int
	for i, e := range nums {
		if !e {
			result = i
			break
		}
	}

	fmt.Println(result)
}
