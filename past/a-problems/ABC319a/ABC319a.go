package main

import "fmt"

func main() {
	var S string
	fmt.Scan(&S)

	/*
		tourist 3858
		ksun48 3679
		Benq 3658
		Um_nik 3648
		apiad 3638
		Stonefeang 3630
		ecnerwala 3613
		mnbvmar 3555
		newbiedmy 3516
		semiexp 3481
	*/
	if S == "tourist" {
		fmt.Println(3858)
	} else if S == "ksun48" {
		fmt.Print(3679)
	} else if S == "Benq" {
		fmt.Print(3658)
	} else if S == "Um_nik" {
		fmt.Print(3648)
	} else if S == "apiad" {
		fmt.Print(3638)
	} else if S == "Stonefeang" {
		fmt.Print(3630)
	} else if S == "ecnerwala" {
		fmt.Print(3613)
	} else if S == "mnbvmar" {
		fmt.Print(3555)
	} else if S == "newbiedmy" {
		fmt.Print(3516)
	} else if S == "semiexp" {
		fmt.Print(3481)
	}
}
