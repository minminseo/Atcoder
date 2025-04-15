package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	N, _ := strconv.Atoi(sc.Text())

	persons := make([]int, N)
	for i := 0; i < N; i++ {
		sc.Scan()
		val, _ := strconv.Atoi(sc.Text())
		persons[i] = val
	}

	bibNumber := make([]int, N)
	for i := 0; i < N; i++ {
		sc.Scan()
		val, _ := strconv.Atoi(sc.Text())
		bibNumber[i] = val
	}

	result := make([]int, N)
	for i := 0; i < N; i++ {
		result[bibNumber[i]-1] = bibNumber[persons[i]-1]
	} // resultのbibNumber[i]-1の位置にbibNumber[i]が見つめてる人の識別番号を入れる。これで埋めていけば、ゼッケン番号の昇順で並んでいる人の見つめ先の人のゼッケンを順に並べられる。

	for i, v := range result {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
	fmt.Println()
}
