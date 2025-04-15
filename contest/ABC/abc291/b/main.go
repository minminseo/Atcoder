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
	scanner.Scan()
	parts := strings.Split(scanner.Text(), " ")

	numbers := make([]int, len(parts))
	for i, part := range parts {
		num, _ := strconv.Atoi(part)
		numbers[i] = num
	}

	sort.Ints(numbers)

	total := 0
	for i := n; i < len(numbers)-n; i++ {
		total += numbers[i]
	}

	count := len(numbers) - 2*n

	average := float64(total) / float64(count)
	fmt.Println(average)
}
