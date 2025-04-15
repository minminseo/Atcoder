package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	line, _ := reader.ReadString('\n')
	parts := strings.Fields(line)
	N, _ := strconv.Atoi(parts[0])
	M, _ := strconv.Atoi(parts[1])

	line, _ = reader.ReadString('\n')
	parts = strings.Fields(line)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		A[i], _ = strconv.Atoi(parts[i])
	}

	line, _ = reader.ReadString('\n')
	parts = strings.Fields(line)
	B := make([]int, M)
	for i := 0; i < M; i++ {
		B[i], _ = strconv.Atoi(parts[i])
	}

	sort.Sort(sort.Reverse(sort.IntSlice(A)))
	sort.Sort(sort.Reverse(sort.IntSlice(B)))

	S := make([]int, N+1)
	T := make([]int, M+1)
	maxT := make([]int, M+1)

	for i := 0; i < N; i++ {
		S[i+1] = S[i] + A[i]
	}

	for i := 0; i < M; i++ {
		T[i+1] = T[i] + B[i]
		maxT[i+1] = max(maxT[i], T[i+1])
	}

	ans := 0
	for i := 0; i <= N; i++ {
		ans = max(ans, S[i]+maxT[min(i, M)])
	}

	fmt.Println(ans)
}
