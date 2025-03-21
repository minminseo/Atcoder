package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	S := sc.Text()

	count := make(map[rune]int)
	for _, r := range S {
		count[r]++
	}

	freq := make(map[int]int)
	for _, c := range count {
		freq[c]++
	}

	for _, v := range freq {
		if v != 0 && v != 2 {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
