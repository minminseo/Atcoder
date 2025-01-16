package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}
	for i := 0; i < n; i++ {
		b[i] = nextInt()
	}
	count1, count2 := 0, 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j && a[i] == b[j] {
				count1++
			} else if a[i] == b[j] {
				count2++
			}
		}
	}
	fmt.Println(count1)
	fmt.Println(count2)
}
