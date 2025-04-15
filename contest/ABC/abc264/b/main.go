package main

import "fmt"

func main() {
	var r, c int
	fmt.Scan(&r, &c)

	d1 := 8 - r
	if d1 < 0 {
		d1 = -d1
	}
	d2 := 8 - c
	if d2 < 0 {
		d2 = -d2
	}
	d := d1
	if d2 > d {
		d = d2
	}

	if d%2 == 0 {
		fmt.Println("white")
	} else {
		fmt.Println("black")
	}
}
