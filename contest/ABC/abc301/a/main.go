package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var s string
	fmt.Scan(&s)

	countA := 0
	countT := 0
	for i := 0; i < n; i++ {
		if s[i] == 'A' {
			countA++
		} else {
			countT++
		}
	}

	if countA > countT {
		fmt.Println("A")
	} else if countA < countT {
		fmt.Println("T")
	} else {
		goalCountA := 0
		goalCountT := 0
		for i := 0; i < n; i++ {
			if s[i] == 'A' {
				goalCountA++
			} else {
				goalCountT++
			}
			if goalCountA == countA {
				fmt.Println("A")
				break
			} else if goalCountT == countT {
				fmt.Println("T")
				break
			}
		}
	}
}
