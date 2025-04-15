package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var s, t string
	fmt.Fscan(in, &s, &t)

	piv := (rune(t[0]) - rune(s[0]) + 26) % 26

	for i := 0; i < len(s); i++ {
		if piv != (rune(t[i])-rune(s[i])+26)%26 {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
