package main

import "fmt"

func main() {
	var S string
	fmt.Scanf("%s", &S)

	for i := 0; i < len(S); i++ {
		fmt.Print(string(S[i]), " ")
	}
}
