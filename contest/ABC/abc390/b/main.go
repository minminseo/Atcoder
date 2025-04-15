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

	line, _ = reader.ReadString('\n')
	part = strings.Fields(line)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		A[i], _ = strconv.Atoi(part[i])
	}

	for i := 1; i < N-1; i++ {
		if A[i]*A[i] != A[i-1]*A[i+1] {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
