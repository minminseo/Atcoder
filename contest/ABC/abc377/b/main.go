package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	br := bufio.NewReader(os.Stdin)

	x := [8][8]bool{}
	for i := 0; i < 8; i++ {
		line, _ := br.ReadString('\n')
		line = strings.TrimSpace(line)
		for j, v := range line {
			if v == '#' {
				for z := 0; z < 8; z++ {
					x[i][z] = true
					x[z][j] = true
				}
			}
		}
	}

	cc := 0
	for i := range x {
		for j := range x[i] {
			if !x[i][j] {
				cc++
			}
		}
	}

	fmt.Println(cc)
}
