package main

import "fmt"

var (
	N int
	L int
	A int
)

func main() {
	fmt.Scan(&N, &L)
	counter := 0
	for i := 0; i < N; i++ {
		fmt.Scan(&A)
		if A >= L {
			counter++
		}
	}
	fmt.Println(counter)
}
