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
		if !seenLeft[slice[i]] { // スライスAの各数値をキーに持ち。そのキーが存在するかどうかかを判定。
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
		sum := left[i] + right[i+1] // 左側の要素i個目以下＋右側の要素i+1個目以上を足す
		if sum > maxSum {
			maxSum = sum
		}
	}

	fmt.Println(maxSum)
}
