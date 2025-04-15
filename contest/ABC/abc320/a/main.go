package main

import (
	"fmt"
	"math"
)

func main() {
	var A, B float64
	fmt.Scan(&A, &B)

	fmt.Println(int(math.Pow(A, B)) + int(math.Pow(B, A)))
}
