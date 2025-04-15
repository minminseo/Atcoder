package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var s string
	fmt.Fscan(reader, &s)

	runes := []rune(s)
	n := len(runes)
	cur := 0

	for cur < n-1 {
		if runes[cur] == 'W' && runes[cur+1] == 'A' {
			runes[cur] = 'A'
			runes[cur+1] = 'C'
			if cur > 0 {
				cur-- // 置換したら一つ前に戻る
			}
		} else {
			cur++
		}
	}

	fmt.Println(string(runes))
}
