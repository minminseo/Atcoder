package main

import "fmt"

func main() {
	N := 0

	fmt.Scanf("%d", &N)

	const bd = 1 << 31

	if N < bd && N >= (-1)*bd {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
