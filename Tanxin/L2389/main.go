package L2389

import "sort"

func answerQueries(nums []int, queries []int) []int {
	sort.Ints(nums)
	for index, v := range queries {
		var sum = 0
		var tlen = 0
		for i := 0; i < len(nums); i++ {
			sum += nums[i]
			if sum > v {
				queries[index] = tlen
				break
			}
			tlen++
			if tlen == len(nums) {
				queries[index] = tlen
				break
			}
		}
	}
	return queries
}
