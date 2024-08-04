package main

import "fmt"

func main() {
	var N string
	fmt.Scan(&N)
	check := true

	for i := 0; i < len(N)-1; i++ {
		if N[i] < N[i+1] {
			check = false
		}
	}
	if check {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
