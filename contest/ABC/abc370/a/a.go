package main

import "fmt"

func main() {
	var l, r int
	fmt.Scan(&l, &r)

	if (l == 1) && (r == 0) {
		fmt.Println("Yes")
	} else if (l == 0) && (r == 1) {
		fmt.Println("No")
	} else {
		fmt.Println("Invalid")
	}
}
