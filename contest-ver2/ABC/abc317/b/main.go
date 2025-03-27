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
	_, _ = reader.ReadString('\n')
	line, _ := reader.ReadString('\n')
	parts := strings.Split(line, " ")

	var nums []int
	for _, p := range parts {
		n, err := strconv.Atoi(strings.TrimSpace(p))
		if err == nil {
			nums = append(nums, n)
		}
	}

	sort.Ints(nums)

	prev := nums[0] - 1
	for _, num := range nums {
		if num != prev+1 {
			fmt.Println(prev + 1)
			return
		}
		prev = num
	}

	fmt.Println("fail")
}
