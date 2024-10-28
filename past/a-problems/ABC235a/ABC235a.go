package main

import (
	"fmt"
	"strconv"
)

func main() {
	var S string
	fmt.Scan(&S)

	ans1, _ := strconv.Atoi(S)
	ans2, _ := strconv.Atoi(S[1:] + S[:1])
	ans3, _ := strconv.Atoi(S[2:] + S[:2])
	fmt.Println(ans1 + ans2 + ans3)
}
