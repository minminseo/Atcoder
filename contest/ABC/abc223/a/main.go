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

	x, _ := strconv.Atoi(sc.Text())

	if x%100 == 0 && x != 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
