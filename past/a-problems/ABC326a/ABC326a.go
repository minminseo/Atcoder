package main

import "fmt"

func main() {
	var X, Y int
	fmt.Scan(&X, &Y)

	if (X > Y && X <= Y+3) || (X < Y && X >= Y-2) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
