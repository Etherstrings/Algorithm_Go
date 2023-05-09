package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxSubArray([]int{1, 2, 3, 1}))
}

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	var sum = math.MinInt
	var cont = 0
	for _, num := range nums {
		cont += num
		if sum >= cont {

		} else {
			sum = cont
		}
		if cont < 0 {
			cont = 0
		}
	}
	return sum
}
