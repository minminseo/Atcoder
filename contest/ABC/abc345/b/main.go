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
	X, _ := strconv.ParseInt(strings.TrimSpace(line), 10, 64)

	if (X+9) < 0 && (X+9)%10 != 0 {
		fmt.Println((X+9)/10 - 1)
	} else {
		fmt.Println((X + 9) / 10)
	}
}
