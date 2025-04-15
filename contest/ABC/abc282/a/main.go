package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())
	Alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	fmt.Println(Alphabet[:k])
}
