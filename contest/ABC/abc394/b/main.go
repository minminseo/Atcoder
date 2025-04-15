package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	s := make([]string, n)

	for i := 0; i < n; i++ {
		scanner.Scan()
		s[i] = scanner.Text()
	}

	sort.Slice(s, func(i, j int) bool {
		return len(s[i]) < len(s[j])
	})

	fmt.Println(strings.Join(s, ""))
}
