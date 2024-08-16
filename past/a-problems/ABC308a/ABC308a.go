package main

import "fmt"

func main() {
	S := make([]int, 7)
	check := true
	for i := 0; i < 7; i++ {
		fmt.Scan(&S[i])
		if !(S[i] >= 100 && S[i] <= 675) || !(S[i]%25 == 0) {
			check = false
		}
	}

	if check {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
