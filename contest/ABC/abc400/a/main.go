package main

import "fmt"

func main() {

	var a int
	fmt.Scan(&a)

	r := 400 % a
	if r == 0 {
		fmt.Println(400 / a)
	} else {
		fmt.Println("-1")
	}
}
