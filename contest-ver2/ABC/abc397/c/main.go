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
	part := strings.Fields(line)
	N, _ := strconv.Atoi(part[0])

	slice := make([]int, N)
	line, _ = reader.ReadString('\n')
	part = strings.Fields(line)
	for i := 0; i < N; i++ {
		slice[i], _ = strconv.Atoi(part[i])
	}

	left := make([]int, N)
	seenLeft := make(map[int]bool)
	count := 0
	for i := 0; i < N; i++ {
		if !seenLeft[slice[i]] {
			seenLeft[slice[i]] = true
			count++
		}
		left[i] = count
	}

	right := make([]int, N)
	seenRight := make(map[int]bool)
	count = 0
	for i := N - 1; i >= 0; i-- {
		if !seenRight[slice[i]] {
			seenRight[slice[i]] = true
			count++
		}
		right[i] = count
	}

	maxSum := 0
	for i := 0; i < N-1; i++ {
		sum := left[i] + right[i+1]
		if sum > maxSum {
			maxSum = sum
		}
	}

	fmt.Println(maxSum)
}
