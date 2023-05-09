package main

func partitionDisjoint(nums []int) int {
	var ans = 0
	var bigleft = -1
	n := len(nums)
	dp := make([]int, n+1)
	dp[n-1] = nums[n-1]
	for i := n - 1; i >= 0; i-- {
		if i == n-1 {
			continue
		}
		dp[i] = min(dp[i+1], nums[i])
	}
	for i := 0; i < n-1; i++ {

		bigleft = max(bigleft, nums[i])
		if (bigleft <= dp[i+1]) {
			ans = i - 0 + 1;
			break
		}
	}
	return ans;
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}