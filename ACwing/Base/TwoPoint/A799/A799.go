package main

import (
	_ "Algorithm/ACwing/Base/Stack"
	"fmt"
)

const N = 100010

var A [N]int
var B [N]int

func main() {
	var n int
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Scan(&A[i])
	}

	left := 1
	ans := -1
	for right := 1; right <= n; right++ {
		B[A[right]]++
		for left <= right && B[A[right]] > 1 {
			B[A[left]]--
			left++
		}
		ans = max(ans, right-left+1)
	}

	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
