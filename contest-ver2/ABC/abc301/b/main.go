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
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	line, _ = reader.ReadString('\n')
	fields := strings.Fields(line)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i], _ = strconv.Atoi(fields[i])
	}

	result := []int{a[0]}
	for i := 1; i < n; i++ {
		prev := a[i-1]
		curr := a[i]
		if prev < curr {
			for j := prev + 1; j < curr; j++ {
				result = append(result, j)
			}
		} else if prev > curr {
			for j := prev - 1; j > curr; j-- {
				result = append(result, j)
			}
		}
		result = append(result, curr)
	}

	for i := 0; i < len(result); i++ {
		fmt.Print(result[i])
		if i != len(result)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
