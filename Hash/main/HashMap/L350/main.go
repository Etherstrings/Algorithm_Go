package main

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2}
	intersect(nums1, nums2)

}

func intersect(nums1 []int, nums2 []int) []int {
	m1 := make(map[int]int)
	m2 := make(map[int]int)
	var temp = 0
	for _, v := range nums1 {
		if p, ok := m1[v]; ok {
			m1[v] = p + 1
		} else {
			m1[v] = 1
		}
	}

	for _, v := range nums2 {
		if p, ok := m2[v]; ok {
			m2[v] = p + 1
		} else {
			m2[v] = 1
		}
	}

	l := []int{}

	for k, v := range m1 {
		if v1, ok := m2[k]; ok {
			//m2中存在m1
			if v > v1 {
				temp = v1
			} else if v < v1 {
				temp = v
			} else {
				temp = v
			}
			//添加k temp次
			for i := 0; i < temp; i++ {
				l = append(l, k)
			}
		}
	}
	ans := make([]int, len(l))
	for i := 0; i < len(l); i++ {
		ans[i] = l[i]
	}
	return ans
}
