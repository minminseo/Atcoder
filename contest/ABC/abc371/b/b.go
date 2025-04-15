package main

import (
	"fmt"
)

func main() {
	var N, M int
	fmt.Scan(&N, &M)

	firstBornSon := make(map[int]bool)

	for i := 0; i < M; i++ {
		var Ai int
		var Bi string
		fmt.Scan(&Ai, &Bi)

		if Bi == "M" {

			if !firstBornSon[Ai] {
				fmt.Println("Yes")
				firstBornSon[Ai] = true
			} else {
				fmt.Println("No")
			}
		} else {

			fmt.Println("No")
		}
	}
}
