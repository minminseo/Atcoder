package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n int
	fmt.Fscan(r, &n)

	if n < 10 {
		fmt.Fprintf(w, "AGC00%d\n", n)
	} else if n < 42 {
		fmt.Fprintf(w, "AGC0%d\n", n)
	} else {
		fmt.Fprintf(w, "AGC0%d\n", n+1)
	}
}
