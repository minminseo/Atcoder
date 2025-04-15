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
	X, _ := strconv.ParseFloat(part[0], 64)

	switch {
	case X >= 38.0:
		fmt.Println(1)
	case X < 37.5:
		fmt.Println(3)
	default:
		fmt.Println(2)
	}
}
