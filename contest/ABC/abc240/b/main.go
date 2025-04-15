package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	inputs := strings.Split(sc.Text(), " ")

	Map := make(map[string]string, len(inputs))
	count := 0

	for i, c := range inputs {
		if Map[inputs[i]] == "" {
			Map[inputs[i]] = c
			count = count + 1
		}
	}

	fmt.Println(count)

}
