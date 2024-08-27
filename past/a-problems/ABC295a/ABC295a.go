package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	w := make([]string, n)

	check := false
	for i := 0; i < n; i++ {
		fmt.Scan(&w[i])
		// and, not, that, the, you
		if (w[i] == "and") || (w[i] == "not") || (w[i] == "that") || (w[i] == "the") || (w[i] == "you") {
			check = true
			break
		}
	}

	if check {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
