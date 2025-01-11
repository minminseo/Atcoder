package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	ar := make([]int, n)
	br := make([]int, n)
	for i := range ar {
		fmt.Scan(&ar[i], &br[i])
	}
	ans := 0.0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			x := ar[i] - ar[j]
			y := br[i] - br[j]
			v := math.Sqrt(float64(x*x) + float64(y*y))
			if ans < v {
				ans = v
			}
		}
	}
	fmt.Println(ans)
}
