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
	L, _ := strconv.Atoi(parts[0])
	R, _ := strconv.Atoi(parts[1])

	S, _ := reader.ReadString('\n')
	S = strings.TrimSpace(S)

	reversed := []rune(S[L-1 : R])
	for i := 0; i < len(reversed)/2; i++ {
		reversed[i], reversed[len(reversed)-1-i] = reversed[len(reversed)-1-i], reversed[i] // 同時に書かないとスワップできない
	}
	S = S[:L-1] + string(reversed) + S[R:]

	fmt.Println(S)
}
