package main

import "fmt"

func main() {
	var Y int
	fmt.Scan(&Y)

	if Y%4 != 0 {
		fmt.Println("365")
	} else if Y%4 == 0 && Y%100 != 0 {
		fmt.Println("366")
	} else if Y%100 == 0 && Y%400 != 0 {
		fmt.Println("365")
	} else if Y%400 == 0 {
		fmt.Println("366")
	}

}
