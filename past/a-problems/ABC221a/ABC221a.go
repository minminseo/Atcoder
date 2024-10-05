package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	abs := strings.Split(sc.Text(), " ")

	a, _ := strconv.ParseFloat(abs[0], 64)
	b, _ := strconv.ParseFloat(abs[1], 64)

	fmt.Println(int(math.Pow(32, a-b)))
}
