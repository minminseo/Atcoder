package main

import "fmt"

func main() {
	var A, B, C, X int
	fmt.Scanf("%d %d %d %d", &A, &B, &C, &X)

	switch {
	case X <= A:
		fmt.Println(1.0)
	case A < X && X <= B:
		fmt.Println(float64(C) / float64(B-A))
	case X > B:
		fmt.Println(0.0)
	}
}
