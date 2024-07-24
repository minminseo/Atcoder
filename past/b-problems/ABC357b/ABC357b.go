package main

import (
	"fmt"
	"strings"
	"unicode"
)

func countUppercase(s string) int {
	count := 0
	for _, r := range s {
		if unicode.IsUpper(r) {
			count++
		}
	}
	return count
}

func countLowercase(s string) int {
	count := 0
	for _, r := range s {
		if unicode.IsLower(r) {
			count++
		}
	}
	return count
}

func main() {
	var S string
	fmt.Scanf("%s", &S)

	if countUppercase(S) > countLowercase(S) {
		upperText := strings.ToUpper(S)
		fmt.Println(upperText)
	} else {
		lowerText := strings.ToLower(S)
		fmt.Println(lowerText)
	}

}
