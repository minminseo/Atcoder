package main

import (
	"fmt"
	"math"
)

func main() {
	var S string
	fmt.Scan(&S)

	pos := make(map[rune]int)
	for i, char := range S {
		pos[char] = i
	}

	totalDistance := 0
	currentPos := pos['A']
	for _, char := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		totalDistance += int(math.Abs(float64(pos[char] - currentPos)))
		currentPos = pos[char]
	}

	fmt.Println(totalDistance)
}
