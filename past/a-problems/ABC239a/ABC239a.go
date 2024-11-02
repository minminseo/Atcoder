package main

import (
	"fmt"
	"math"
)

func main() {
	var H float64
	fmt.Scan(&H)

	fmt.Println(math.Sqrt(H * (12800000 + H)))
}
