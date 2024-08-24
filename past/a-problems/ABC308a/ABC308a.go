package main

import "fmt"

func main() {
    S := make([]int, 8)
    check := true

    for i := 0; i < 8; i++ {
        fmt.Scan(&S[i])
        if !(S[i] >= 100 && S[i] <= 675) || !(S[i]%25 == 0) {
            check = false
        }
    }

    for i := 1; i < 8; i++ {
        if S[i-1] > S[i] {
            check = false
        }
    }

    if check {
        fmt.Println("Yes")
    } else {
        fmt.Println("No")
    }
}
