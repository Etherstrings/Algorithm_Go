package main

func moveZeroes(nums []int) {

	var right = 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[right] = nums[i]
			right++
		}
	}

	for i := right; i < len(nums); i++ {
		nums[i] = 0
	}

}
