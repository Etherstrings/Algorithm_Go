package main

import (
	"fmt"
)

const N = 100010

var (
	A [N]int
	B [N]int
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	if n > m {
		fmt.Println("No")
	} else {
		for i := 1; i <= n; i++ {
			fmt.Scan(&A[i])
		}
		for i := 1; i <= m; i++ {
			fmt.Scan(&B[i])
		}
		slow := 1
		for fast := 1; fast <= m; fast++ {
			if slow == n+1 {
				break
			}
			if B[fast] == A[slow] {
				slow++
			}
		}
		if slow != n+1 {
			fmt.Println("No")
		} else {
			fmt.Println("Yes")
		}
	}
}
