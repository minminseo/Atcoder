package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	part := strings.Fields(line)
	S := part[0]

	i := 0
	insertionCount := 0
	stringI := "i"
	for i < len(S) {
		if string(S[i]) == stringI {
			i++
		} else {
			insertionCount++
		}
		if stringI == "i" {
			stringI = "o"
		} else {
			stringI = "i"
		}
	}
	gattai := len(S) + insertionCount
	if gattai%2 != 0 {
		insertionCount++
	}

	fmt.Println(insertionCount)
}
