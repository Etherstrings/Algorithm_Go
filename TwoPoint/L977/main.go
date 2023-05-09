package main

func sortedSquares(nums []int) []int {
	//双指针
	var left = 0
	var right = len(nums) - 1

	ans := make([]int, len(nums))

	for i := len(nums) - 1; i >= 0; i-- {
		if left > right {
			break
		}
		var l = nums[left] * nums[left]
		var r = nums[right] * nums[right]
		if l >= r {
			ans[i] = l
			left++
		} else if l < r {
			ans[i] = r
			right--
		}
	}
	return ans
}
