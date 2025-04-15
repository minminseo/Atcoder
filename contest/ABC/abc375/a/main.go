package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, _ := reader.ReadString('\n')
	n = strings.TrimSpace(n)
	N, _ := strconv.Atoi(n)
	s, _ := reader.ReadString('\n')

	count := 0
	for i := 1; i < N-1; i++ {
		if s[i-1] == '#' && s[i] == '.' && s[i+1] == '#' {
			count++
		}
	}

	fmt.Println(count)
}
