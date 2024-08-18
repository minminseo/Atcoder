package main

import "fmt"

func main() {
	var A, B, C int
	fmt.Scan(&A, &B, &C)

	if B > C {
		if (A <= B) && (A >= C) {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	} else {
		if ((A <= 24) && (A >= C)) || (A <= B) {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}

	}

}
