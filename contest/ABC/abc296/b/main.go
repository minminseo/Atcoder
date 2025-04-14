package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	board := make([]string, 8)
	for i := 0; i < 8; i++ {
		scanner.Scan()
		board[i] = scanner.Text()
	}

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if board[i][j] == '*' {
				if j == 0 {
					fmt.Print("a")
				} else if j == 1 {
					fmt.Print("b")
				} else if j == 2 {
					fmt.Print("c")
				} else if j == 3 {
					fmt.Print("d")
				} else if j == 4 {
					fmt.Print("e")
				} else if j == 5 {
					fmt.Print("f")
				} else if j == 6 {
					fmt.Print("g")
				} else if j == 7 {
					fmt.Print("h")
				}

				fmt.Println(8 - i)
			}
		}
	}
}
