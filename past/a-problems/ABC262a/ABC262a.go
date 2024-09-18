package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	y, _ := strconv.Atoi(sc.Text())

	if y%4 == 2 {
		fmt.Println(y)
	} else if y%4 < 3 {
		fmt.Println((y - y%4) + 2)
	} else {
		fmt.Println((y - y%4) + 6)
	}
}
