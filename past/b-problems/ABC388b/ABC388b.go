package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		n, d int
	)
	fmt.Scanf("%d %d", &n, &d)

	sc := bufio.NewScanner(os.Stdin)

	input := make([][]int, 0, n)

	for i := 0; i < n; i++ {
		sc.Scan()
		inputs := strings.Split(sc.Text(), " ")
		tmp := make([]int, 0, 2)
		for _, v := range inputs {
			t, _ := strconv.Atoi(v)
			tmp = append(tmp, t)
		}
		input = append(input, tmp)
	}

	for i := 1; i <= d; i++ {
		max := 0
		for _, v := range input {
			now := v[0] * (v[1] + i)
			if now > max {
				max = now
			}
		}
		fmt.Println(max)
	}
}
