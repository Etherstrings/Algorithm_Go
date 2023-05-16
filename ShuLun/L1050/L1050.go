package main

import (
	"fmt"
)

func smallestRepunitDivByK(k int) int {
	if k%5 == 0 || k%2 == 0 {
		return -1
	}
	seen := make(map[int]bool)
	x := 1 % k
	for x > 0 && !seen[x] {
		seen[x] = true
		x = (x*10 + 1) % k
	}
	if x > 0 {
		return -1
	}
	return len(seen) + 1
}

func main() {
	fmt.Println(smallestRepunitDivByK(1))
	fmt.Println(smallestRepunitDivByK(2))
	fmt.Println(smallestRepunitDivByK(3))
	fmt.Println(smallestRepunitDivByK(5))
	fmt.Println(smallestRepunitDivByK(7))
}
