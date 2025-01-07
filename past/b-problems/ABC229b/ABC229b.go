package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var a, b string
	fmt.Scan(&a, &b)
	l := int(math.Min(float64(len(a)), float64(len(b))))
	for i := 1; i <= l; i++ {
		x, _ := strconv.Atoi(string(a[len(a)-i]))
		y, _ := strconv.Atoi(string(b[len(b)-i]))
		if x+y > 9 {
			fmt.Print("Hard")
			return
		}
	}
	fmt.Print("Easy")
}
