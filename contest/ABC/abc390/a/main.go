package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanLines)
	sc.Scan()
	s := sc.Text()
	result := "No"
	if s == "2 1 3 4 5" || s == "1 3 2 4 5" || s == "1 2 4 3 5" || s == "1 2 3 5 4" {
		result = "Yes"
	}
	fmt.Println(result)
}
