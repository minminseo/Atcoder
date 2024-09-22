package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	s := sc.Text()

	if s == "Hello,World!" {
		fmt.Println("AC")
	} else {
		fmt.Println("WA")
	}
}
