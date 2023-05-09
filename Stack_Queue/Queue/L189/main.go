package main

import "Algorithm/Queue"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	var k = 3
	rotate(nums, k)
}

func rotate(nums []int, k int) {
	var que = Queue.NewQueue()
	for i := len(nums) - 1; i >= 0; i-- {
		que.Offer(nums[i])
	}
	for i := 0; i < k; i++ {
		v, _ := que.Poll()
		que.Offer(v)
	}
	for i := len(nums) - 1; i >= 0; i-- {
		v, _ := que.Poll()
		//在这里转换形式
		v1 := v.(int)
		nums[i] = v1
	}

}
