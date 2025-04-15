package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func BinarySearchExist(arr []int, target int) bool {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == target {
			return true
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	parts := strings.Fields(line)
	N, _ := strconv.Atoi(parts[0])
	M, _ := strconv.Atoi(parts[1])

	line, _ = reader.ReadString('\n')
	parts = strings.Fields(line)
	A := make([]int, M)
	for i := 0; i < M; i++ {
		A[i], _ = strconv.Atoi(parts[i])
	}

	sort.Ints(A)

	missing := []int{}
	for i := 1; i <= N; i++ {
		if !BinarySearchExist(A, i) {
			missing = append(missing, i)
		}
	}

	fmt.Println(len(missing))

	if len(missing) > 0 {
		for i, num := range missing {
			if i > 0 {
				fmt.Print(" ") // fmt.Print(num, " ")これだと行末にスペース入って誤答扱いになる（実際はこれでもAC出る）
			}
			fmt.Print(num)
		}
		fmt.Println()
	} else {
		fmt.Println()
	}

}
