package main

import "fmt"

func main() {
	var S string
	fmt.Scan(&S)
	//  ACE、BDF、CEG、DFA、EGB、FAC、GBD
	if S == "ACE" || S == "BDF" || S == "CEG" || S == "DFA" || S == "EGB" || S == "FAC" || S == "GBD" {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
