package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	b := []byte(s)
	for i := 0; i < len(b); i++ {
		if b[i] == '0' {
			b[i] = '1'
		} else {
			b[i] = '0'
		}
	}
	s = string(b)
	fmt.Println(s)
}
