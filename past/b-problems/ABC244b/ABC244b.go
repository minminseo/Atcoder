package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

const (
	initialBufSize = 10000
)

func main() {
	buf := make([]byte, initialBufSize)
	sc.Buffer(buf, 9223372036854775807)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	t := sc.Text()
	IsN := false
	IsE := true
	IsS := false
	IsW := false
	x := 0
	y := 0
	for i := 0; i < n; i++ {
		if t[i:i+1] == "R" {
			if IsE {
				IsE = false
				IsS = true
			} else if IsS {
				IsS = false
				IsW = true
			} else if IsW {
				IsW = false
				IsN = true
			} else if IsN {
				IsN = false
				IsE = true
			}
		} else {
			if IsE {
				x++
			} else if IsS {
				y--
			} else if IsW {
				x--
			} else if IsN {
				y++
			}
		}
	}
	fmt.Println(strconv.Itoa(x) + " " + strconv.Itoa(y))
}
