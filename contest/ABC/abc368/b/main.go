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
	line, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	line, _ = reader.ReadString('\n')
	fields := strings.Fields(line)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i], _ = strconv.Atoi(fields[i])
	}

	steps := 0
	for {
		count := 0
		for _, v := range a {
			if v > 0 {
				count++
			}
		}
		if count <= 1 {
			break
		}

		sort.Sort(sort.Reverse(sort.IntSlice(a)))
		a[0]--
		a[1]--
		steps++
	}

	fmt.Println(steps)
}
