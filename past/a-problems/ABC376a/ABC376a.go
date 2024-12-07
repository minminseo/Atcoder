package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var N, C int
	fmt.Scanf("%d %d", &N, &C)
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	Ts := strings.Split(strings.TrimSpace(input), " ")
	t, _ := strconv.Atoi(Ts[0])
	count := 1
	for i := 1; i < N; i++ {
		cur, _ := strconv.Atoi(Ts[i])
		if cur-t >= C {
			t = cur
			count++
		}
	}
	fmt.Println(count)
}
