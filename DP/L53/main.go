package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxSubArray([]int{1, 2, 3, 1}))
}

func maxSubArray(nums []int) int {
	//dp[i]：包括下标i之前的最大连续子序列和为dp[i]。
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 0 {
		return 0
	}

	var dp = make([]int, len(nums))

	dp[0] = nums[0]

	for i := 1; i < len(dp); i++ {
		if (dp[i-1]+nums[i] >= nums[i]) {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
	}

	var ans = math.MinInt
	for _, v := range dp {
		if (v >= ans) {
			ans = v
		} else {

		}
	}
	return ans
}
