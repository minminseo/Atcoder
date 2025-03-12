package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	parts := strings.Split(scanner.Text(), " ")
	_, _ = strconv.Atoi(parts[0])
	M, _ := strconv.Atoi(parts[1])

	ans := 0
	edgeMap := make(map[[2]int]int)

	for i := 0; i < M; i++ {
		scanner.Scan()
		line := strings.Split(scanner.Text(), " ")
		u, _ := strconv.Atoi(line[0])
		v, _ := strconv.Atoi(line[1])

		if u == v {
			ans++
			continue
		}
		if u > v {
			u, v = v, u
		}
		key := [2]int{u, v}
		edgeMap[key]++ // edgeMapに頂点のペアをキーとして保存し、valueは多重辺じゃないなら1、多重辺なら2以上になる。
	}

	for _, count := range edgeMap {
		ans += count - 1 // edgeMapの各valueをcountとして、多重辺じゃない場合1のためcount-1は0となり加算されない。
	}

	fmt.Println(ans)
}
