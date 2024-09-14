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
	a, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	b, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	c, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	d, _ := strconv.Atoi(scanner.Text())

	calculation := (a + b) * (c - d)
	fmt.Println(calculation)
	fmt.Println("Takahashi")
}
