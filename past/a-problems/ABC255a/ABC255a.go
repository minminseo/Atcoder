package main

import "fmt"

func main() {
	var R, C, A11, A12, A21, A22 int
	fmt.Scanf("%d %d", &R, &C)
	fmt.Scanf("%d %d", &A11, &A12)
	fmt.Scanf("%d %d", &A21, &A22)

	if R == 1 {
		if C == 1 {
			fmt.Println(A11)
		} else {
			fmt.Println(A12)
		}
	} else {
		if C == 1 {
			fmt.Println(A21)
		} else {
			fmt.Println(A22)
		}
	}
}
