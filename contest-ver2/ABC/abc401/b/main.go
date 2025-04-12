package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	isLoggedIn := false
	authErrors := 0

	for i := 0; i < N; i++ {
		var S string
		fmt.Scan(&S)

		switch S {
		case "login":
			isLoggedIn = true
		case "logout":
			isLoggedIn = false
		case "private":
			if !isLoggedIn {
				authErrors++
			}
		}
	}

	fmt.Println(authErrors)
}
