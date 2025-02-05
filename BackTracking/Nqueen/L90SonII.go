package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("answer")
}

// SubsetsWithDup 返回所有可能的子集，包含重复元素但不包含重复的子集
func subsetsWithDup(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	used := make([]bool, len(nums))
	sort.Ints(nums)
	var ans [][]int
	var path []int
	Backtracking(nums, used, 0, &ans, &path)
	return ans
}

func Backtracking(nums []int, used []bool, startIndex int, ans *[][]int, path *[]int) {
	// 将当前路径加入结果集中
	*ans = append(*ans, append([]int(nil), *path...))

	if startIndex >= len(nums) {
		return
	}

	for i := startIndex; i < len(nums); i++ {
		// 跳过重复元素
		if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
			continue
		}
		// 选择当前元素
		*path = append(*path, nums[i])
		used[i] = true
		Backtracking(nums, used, i+1, ans, path)
		// 撤销选择
		used[i] = false
		*path = (*path)[:len(*path)-1]
	}
}
