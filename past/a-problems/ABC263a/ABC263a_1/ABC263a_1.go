package main

import (
	"fmt"
)

func main() {
	var a, b, c, d, e string
	fmt.Scan(&a, &b, &c, &d, &e)
	abcdes := []string{a, b, c, d, e}

	first := ""
	second := ""
	for i := 0; i < 5; i++ {
		if i == 0 {
			first = abcdes[0]
		}
		if first != abcdes[i] {
			second = abcdes[i]
			break
		}
	}

	countFirst := 0
	countSecond := 0
	for i := 0; i < 5; i++ {
		if abcdes[i] == first {
			countFirst++
		} else if abcdes[i] == second {
			countSecond++
		}
	}

	if (countFirst == 3 && countSecond == 2) || (countFirst == 2 && countSecond == 3) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
