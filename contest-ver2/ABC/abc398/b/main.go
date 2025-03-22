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

	count := make(map[int]int)
	for _, p := range parts {
		n, _ := strconv.Atoi(p)
		count[n]++
	}

	values := []int{}
	for _, c := range count {
		values = append(values, c)
	}

	fullHouse := false
	for i := 0; i < len(values); i++ {
		for j := 0; j < len(values); j++ {
			if i == j {
				continue // 被ってる位置以外の全部の組み合わせを調べる
			}
			if values[i] >= 3 && values[j] >= 2 {
				fullHouse = true
			}
		}
	}

	if fullHouse {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
