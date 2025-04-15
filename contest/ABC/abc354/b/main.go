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
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	nStr := strings.TrimSpace(line)
	n, _ := strconv.Atoi(nStr)

	users := make([]string, n)
	rates := make(map[string]int)

	for i := 0; i < n; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		parts := strings.Fields(line)

		name := parts[0]
		rate, _ := strconv.Atoi(parts[1])

		users[i] = name
		rates[name] = rate
	}

	sort.Strings(users)

	total := 0
	for _, name := range users {
		total += rates[name]
	}

	winnerIndex := total % n
	fmt.Println(users[winnerIndex])
}
