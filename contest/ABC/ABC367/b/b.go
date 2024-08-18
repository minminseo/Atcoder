package main

import "fmt"

func main() {
	var x string
	fmt.Scan(&x)

	// array := []rune(x)

	count := 0
	for i := len(x) - 1; i >= 0; i-- {
		if string(x[i]) == "0" {
			count++
		} else {
			break
		}
	}
	if count == 3 {
		fmt.Println(string(x[0 : len(x)-count-1]))
	} else {
		fmt.Println(x[0 : len(x)-count])
	}
}
