package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	sc.Scan()
	a := strings.Split(sc.Text(), " ")

	S1, S2 := a[0], a[1]

	if S1 == S2 {
		if S1 == "sick" {
			fmt.Println(1)
		} else {
			fmt.Println(4)
		}
	} else {
		if S1 == "sick" {
			fmt.Println(2)
		} else {
			fmt.Println(3)
		}
	}
}
