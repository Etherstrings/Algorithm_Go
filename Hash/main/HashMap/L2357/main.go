package main

func minimumOperations(nums []int) int {
	m := make(map[int]int)

	for _, v := range nums {
		if p, ok := m[v]; ok {
			m[v] = p + 1
		} else {
			m[v] = 1
		}
	}
	var size = 0
	var flag bool = false
	for key, _ := range m {
		size++
		if key == 0 {
			flag = true
		}
	}
	if size == 1 && flag == true {
		return 0
	}
	if flag == true {
		return size - 1
	}
	return size
}
