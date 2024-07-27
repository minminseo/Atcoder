package main

import "fmt"

func check(resultList []string) bool {
	for i := 0; i < len(resultList)-2; i++ {
		if resultList[i] == "sweet" && resultList[i+1] == "sweet" {
			return true
		}
	}
	return false
}

func main() {
	var N int
	resultList := make([]string, N)
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		var S string
		fmt.Scan(&S)
		resultList = append(resultList, S)
	}

	if check(resultList) {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}
