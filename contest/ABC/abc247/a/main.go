package main

import "fmt"

func main() {
	var S string
	fmt.Scan(&S)
	fmt.Println("0" + S[:3])
}
