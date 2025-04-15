package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	stack := make([]int, 100)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	Q, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < Q; i++ {
		scanner.Scan()
		query := strings.Split(scanner.Text(), " ")

		if query[0] == "1" {
			x, _ := strconv.Atoi(query[1])
			stack = append(stack, x)
		} else {
			if len(stack) > 0 {
				fmt.Println(stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
		}
	}
}
