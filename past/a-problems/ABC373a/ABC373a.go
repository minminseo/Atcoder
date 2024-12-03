package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	num := 0
	index := 0
	for scanner.Scan() {
		index++
		line := scanner.Text()
		if line == "" {
			break
		}
		if index == len(line) {
			num++
		}
	}
	fmt.Println(num)
}
