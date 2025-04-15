package main

import "fmt"

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)

	check := true
	if len(s) == 1 {
		fmt.Println("Yes")
	} else {
		for i := 0; i < n-1; i++ {
			if (s[i] == 'M') && (s[i+1] == 'F') {
				check = true
			} else if (s[i] == 'F') && (s[i+1] == 'M') {
				check = true
			} else {
				check = false
				break
			}
		}
		if check {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
