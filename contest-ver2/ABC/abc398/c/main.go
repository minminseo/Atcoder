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
	line, _ := reader.ReadString('\n')
	parts := strings.Fields(line)
	N, _ := strconv.Atoi(parts[0])

	line, _ = reader.ReadString('\n')
	parts = strings.Fields(line)

	A := make([]int, N)
	count := make(map[int]int)

	for i := 0; i < N; i++ {
		A[i], _ = strconv.Atoi(parts[i])
		count[A[i]]++
	}

	max := -1
	// check := false
	// result := 0
	index := -1
	for i := 0; i < N; i++ {
		if count[A[i]] == 1 {
			if A[i] > max {
				max = A[i]
				// check = true
				// result = i
				index = i
			}
		}
	}

	if index != -1 {
		fmt.Println(index + 1)
	} else {
		fmt.Println(index)
	}
}
