package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	str := "abcdefghijklmnopqrstuvwxyz"
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	list := strings.Split(sc.Text(), " ")
	var ans string
	for i := 0; i < 26; i++ {
		index, _ := strconv.Atoi(list[i])
		ans += string(str[index-1])
	}
	fmt.Println(ans)
}
