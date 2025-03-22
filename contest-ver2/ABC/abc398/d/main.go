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
	parts := strings.Fields(line)
	N, _ := strconv.Atoi(parts[0])
	R, _ := strconv.Atoi(parts[1])
	C, _ := strconv.Atoi(parts[2])

	Sline, _ := reader.ReadString('\n')
	S := strings.TrimSpace(Sline)

	posKey := func(r, c int) string {
		return strconv.Itoa(r) + "," + strconv.Itoa(c)
	}

	fR, fC := 0, 0
	smoke := make(map[string]bool)
	smoke[posKey(fR, fC)] = true

	result := make([]byte, N)

	for i := 0; i < N; i++ {
		// 煙の更新数は最悪等差数列なのでO(N^２)→焚き火と高橋くんを逆方向に移動させて、煙を時刻毎に毎回更新するのを防ぐことでO(N)にする。
		switch S[i] {
		case 'N':
			R++
			fR++
		case 'W':
			C++
			fC++
		case 'S':
			R--
			fR--
		case 'E':
			C--
			fC--
		}

		smoke[posKey(fR, fC)] = true

		if smoke[posKey(R, C)] {
			result[i] = '1'
		} else {
			result[i] = '0'
		}
	}

	fmt.Println(string(result))
}
