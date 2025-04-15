package main

import "fmt"

func main() {
	var A, B, C, D, E, F, X int
	fmt.Scanf("%d %d %d %d %d %d %d", &A, &B, &C, &D, &E, &F, &X)

	var takahashi, aoki int
	for i := 0; i < X; i++ {
		if i%(A+C) < A {
			takahashi += B
		}

		if i%(D+F) < D {
			aoki += E
		}
	}

	if takahashi == aoki {
		fmt.Println("Draw")
	} else if takahashi > aoki {
		fmt.Println("Takahashi")
	} else {
		fmt.Println("Aoki")
	}
}
