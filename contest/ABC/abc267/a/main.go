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
	s := scanner.Text()

	//  Monday, Tuesday, Wednesday, Thursday, Friday
	if s == "Monday" {
		fmt.Println(5)
	} else if s == "Tuesday" {
		fmt.Println(4)
	} else if s == "Wednesday" {
		fmt.Println(3)
	} else if s == "Thursday" {
		fmt.Println(2)
	} else if s == "Friday" {
		fmt.Println(1)
	}
}
