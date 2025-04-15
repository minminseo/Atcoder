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
	hw := strings.Fields(line)
	H, _ := strconv.Atoi(hw[0])
	W, _ := strconv.Atoi(hw[1])

	line, _ = reader.ReadString('\n')
	sij := strings.Fields(line)
	sI, _ := strconv.Atoi(sij[0])
	sJ, _ := strconv.Atoi(sij[1])
	sI--
	sJ--

	C := make([]string, H)
	for i := 0; i < H; i++ {
		line, _ = reader.ReadString('\n')
		C[i] = strings.TrimSpace(line)
	}

	line, _ = reader.ReadString('\n')
	X := strings.TrimSpace(line)
	for _, o := range X {
		ni, nj := sI, sJ // 高橋くんの位置の更新用

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

		// Xの各文字に対して、グリッド範囲内かと移動先があいているかどうか確認（"."か"#"か）
		if ni >= 0 && ni < H && nj >= 0 && nj < W && C[ni][nj] == '.' {
			sI, sJ = ni, nj // 高橋くんの位置更新
		}
	}

	fmt.Println(sI+1, sJ+1)
}
