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
	xy := sc.Text()

	parts := strings.Split(xy, ".")
	result := parts[1]

	finalResult, _ := strconv.Atoi((result))
	if 0 <= finalResult && finalResult <= 2 {
		fmt.Println(parts[0] + "-")
	} else if 3 <= finalResult && finalResult <= 6 {
		fmt.Println(parts[0])
	} else {
		fmt.Println(parts[0] + "+")
	}
}
