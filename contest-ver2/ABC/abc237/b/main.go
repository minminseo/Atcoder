package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	parts := strings.Fields(sc.Text())
	H, _ := strconv.Atoi(parts[0])
	W, _ := strconv.Atoi(parts[1])

	A := make([][]int, H)
	B := make([][]int, W)
	for i := range B {
		B[i] = make([]int, H)
	}

	for i := 0; i < H; i++ {
		sc.Scan()
		row := strings.Fields(sc.Text())
		A[i] = make([]int, W)
		for j := 0; j < W; j++ {
			val, _ := strconv.Atoi(row[j])
			A[i][j] = val
			B[j][i] = val // iとjを逆にして転置行列を同時に代入する
		}
	}

	for i := 0; i < W; i++ {
		for j := 0; j < H; j++ { // 転置行列の出力のためにHとWを逆にして出力する
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(B[i][j])
		}
		fmt.Println()
	}
}
