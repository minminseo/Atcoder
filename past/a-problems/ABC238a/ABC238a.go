package main

import (
	"fmt"
	"math"
)

func main() {
	n := 0.0
	fmt.Scan(&n)

	tn := math.Pow(2, n)
	nt := math.Pow(n, 2)

	if tn > nt {
		fmt.Print("Yes")
	} else {
		fmt.Print("No")
	}
}
