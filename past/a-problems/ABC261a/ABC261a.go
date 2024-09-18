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
	abs := strings.Split(sc.Text(), " ")
	a, _ := strconv.Atoi(abs[0])
	b, _ := strconv.Atoi(abs[1])

	if 0 < a && b == 0 {
		fmt.Println("Gold")
	} else if a == 0 && 0 < b {
		fmt.Println("Silver")
	} else {
		fmt.Println("Alloy")
	}

}
