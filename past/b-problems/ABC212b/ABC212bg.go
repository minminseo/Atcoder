package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	var x = make([]int, 4)
	for i, e := range s {
		x[i] = int(e - '0')
	}
	var samecnt int = 1
	for i := 1; i < 4; i++ {
		if x[0] == x[i] {
			samecnt++
		}
	}
	if samecnt == 4 {
		fmt.Println("Weak")
		return
	}
	var ok = false
	for i := 0; i < 3; i++ {
		n := (x[i] + 1) % 10
		if x[i+1] != n {
			ok = true
			break
		}
	}
	if ok {
		fmt.Println("Strong")
	} else {
		fmt.Println("Weak")
	}
}
