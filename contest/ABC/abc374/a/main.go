package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input := scanner.Text()
		if strings.HasSuffix(input, "san") {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
