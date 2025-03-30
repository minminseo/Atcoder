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
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	line, _ = reader.ReadString('\n')
	parts := strings.Fields(line)

	scores := make([]int, n)
	for i := 0; i < n; i++ {
		scores[i], _ = strconv.Atoi(parts[i])
	}

	indices := make([]int, n)
	for i := 0; i < n; i++ {
		indices[i] = i
	}

	sort.Slice(indices, func(i, j int) bool {
		return scores[indices[i]] > scores[indices[j]]
	})

	ranks := make([]int, n)
	r := 1
	for i := 0; i < n; {
		j := i
		for j < n && scores[indices[j]] == scores[indices[i]] {
			j++
		}
		for k := i; k < j; k++ {
			ranks[indices[k]] = r
		}
		r += j - i
		i = j
	}

	for i := 0; i < n; i++ {
		fmt.Println(ranks[i])
	}
}
