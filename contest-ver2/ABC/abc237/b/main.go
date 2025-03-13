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
	H, _ := strconv.Atoi(parts[0])
	W, _ := strconv.Atoi(parts[1])

	A := make([][]int, H)
	B := make([][]int, W)
	for i := range B {
		B[i] = make([]int, H)
	}

	for i := 0; i < H; i++ {
		line, _ := reader.ReadString('\n')
		row := strings.Fields(line)
		A[i] = make([]int, W)
		for j := 0; j < W; j++ {
			val, _ := strconv.Atoi(row[j])
			A[i][j] = val
			B[j][i] = val // iとjを逆にして転置行列を同時に代入する
		}
	}

	for i := 0; i < W; i++ { // 転置行列の出力のためにHとWを逆にして出力する
		for j := 0; j < H; j++ {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(B[i][j])
		}
		fmt.Println()
	}
}
