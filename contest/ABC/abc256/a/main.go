package main

import (
	"fmt"
	"math"
)

func main() {
	var N float64
	fmt.Scan(&N)

	fmt.Println(int(math.Pow(2, N)))
}
