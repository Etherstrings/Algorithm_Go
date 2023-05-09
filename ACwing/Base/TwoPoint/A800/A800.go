package main

import (
	"fmt"
)

const N = 100010

var (
	A [N]int64
	B [N]int64
)

func main() {
	var n, m, target int
	fmt.Scan(&n, &m, &target)
	for i := 1; i <= n; i++ {
		fmt.Scan(&A[i])
	}
	for i := 1; i <= m; i++ {
		fmt.Scan(&B[i])
	}
	left := 1
	right := m
	for left <= n && right >= 1 {
		if A[left]+B[right] == int64(target) {
			fmt.Println(left-1, right-1)
			break
		} else if A[left]+B[right] > int64(target) {
			right--
		} else {
			left++
		}
	}
}
