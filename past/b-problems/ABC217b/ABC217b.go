package main

import (
	"fmt"
)

func main() {
	var S1, S2, S3 string
	fmt.Scanf("%s", &S1)
	fmt.Scanf("%s", &S2)
	fmt.Scanf("%s", &S3)

	slice := []string{S1, S2, S3}
	list := []string{"ABC", "ARC", "AGC", "AHC"}

	for _, v := range list {
		var ok bool
		for _, v2 := range slice {
			if v == v2 {
				ok = true
			}
		}

		if !ok {
			fmt.Println(v)
			return
		}
	}

}
