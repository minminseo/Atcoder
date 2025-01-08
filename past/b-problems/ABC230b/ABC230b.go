package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Scan()
	S := sc.Text()
	T := "oxxoxxoxxoxxoxx"
	L := len(S)
	d := 0
	for i := 0; i < 15-L; i++ {
		if string(T[i:i+L]) == S {
			fmt.Println("Yes")
			d = 1
			break
		}
	}
	if d == 0 {
		fmt.Print("No")
	}
}
