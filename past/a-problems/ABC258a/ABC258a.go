package main

import (
	"fmt"
)

func main() {
	k := 0
	fmt.Scan(&k)

	hour := 21 + k/60
	minuts := k % 60

	fmt.Printf("%d:%02d", hour, minuts)
}
