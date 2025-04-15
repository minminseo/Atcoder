package main

import (
	"fmt"
	"strings"
)

func main() {
	var S string
	fmt.Scan(&S)
	result := make([]string, 0)

	for i := 0; i < len(S); i++ {
		if S[i] != 'a' && S[i] != 'i' && S[i] != 'u' && S[i] != 'e' && S[i] != 'o' {
			result = append(result, string(S[i]))

		}
	}
	fmt.Println(strings.Join(result, ""))
}
