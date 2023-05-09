package main

import "fmt"

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
}

func containsDuplicate(nums []int) bool {
	var flag bool = false

	var m = make(map[int]int)

	for _, i := range nums {
		if _, ok := m[i]; ok {
			m[i] = m[i] + 1
		} else {
			m[i] = 1
		}
	}
	for _, v := range m {
		if v >= 2 {
			flag = true
			break
		}
	}
	return flag
}
