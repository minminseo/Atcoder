package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	S := strings.Split(sc.Text(), "")
	N := len(S)
	result := 0
	for i := 0; i < N-2; i++ {
		if S[i] != "A" {
			continue
		}
		for j := i + 1; j < N-1; j++ {
			if S[j] != "B" {
				continue
			}
			var k int
			k = 2*j - i
			if k >= N {
				continue
			}
			if S[k] == "C" {
				result += 1
			}
		}
	}
	fmt.Println(result)
}
