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
	hwxy := strings.Fields(line)
	h, _ := strconv.Atoi(hwxy[0])
	w, _ := strconv.Atoi(hwxy[1])
	x, _ := strconv.Atoi(hwxy[2])
	y, _ := strconv.Atoi(hwxy[3])
	x--
	y--

	S := make([][]byte, h)
	for i := 0; i < h; i++ {
		line, _ := reader.ReadString('\n')
		S[i] = []byte(strings.TrimSpace(line))
	}

	line, _ = reader.ReadString('\n')
	T := []byte(strings.TrimSpace(line))

	houses := 0
	for _, o := range T {
		ni, nj := x, y
		switch o {
		case 'L':
			nj--
		case 'R':
			nj++
		case 'U':
			ni--
		case 'D':
			ni++
		}
		if ni > 0 && ni < h && nj > 0 && nj < w && S[ni][nj] != '#' {
			x, y = ni, nj
		}

		if S[x][y] == '@' {
			houses++
			S[x][y] = '.'
		}
	}
	fmt.Println(x+1, y+1, houses)
}
