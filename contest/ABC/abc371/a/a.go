package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	Sab := scanner.Text()
	scanner.Scan()
	Sac := scanner.Text()
	scanner.Scan()
	Sbc := scanner.Text()

	if Sab == "<" {
		if Sac == "<" {
			if Sbc == "<" {
				fmt.Println("B")
			} else {
				fmt.Println("C")
			}
		} else {
			fmt.Println("A")
		}
	} else {
		if Sac == "<" {
			fmt.Println("A")
		} else {
			if Sbc == "<" {
				fmt.Println("C")
			} else {
				fmt.Println("B")
			}
		}
	}

}
